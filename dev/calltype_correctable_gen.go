// DO NOT EDIT. Generated by 'gorums' plugin for protoc-gen-go
// Source file to edit is: calltype_correctable_tmpl

package dev

import (
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"golang.org/x/net/context"
)

/* Methods on Configuration and the correctable struct ReadCorrectableReply */

// ReadCorrectableReply is a reference to a correctable ReadCorrectable quorum call.
type ReadCorrectableReply struct {
	sync.Mutex
	// the actual reply
	*State
	NodeIDs  []uint32
	level    int
	err      error
	done     bool
	watchers []*struct {
		level int
		ch    chan struct{}
	}
	donech chan struct{}
}

// ReadCorrectable asynchronously invokes a
// correctable ReadCorrectable quorum call on configuration c and returns a
// ReadCorrectableReply which can be used to inspect any replies or errors
// when available.
func (c *Configuration) ReadCorrectable(ctx context.Context, args *ReadRequest) *ReadCorrectableReply {
	corr := &ReadCorrectableReply{
		level:   LevelNotSet,
		NodeIDs: make([]uint32, 0, c.n),
		donech:  make(chan struct{}),
	}
	go func() {
		c.mgr.readCorrectable(ctx, c, corr, args)
	}()
	return corr
}

// Get returns the reply, level and any error associated with the
// ReadCorrectable. The method does not block until a (possibly
// itermidiate) reply or error is available. Level is set to LevelNotSet if no
// reply has yet been received. The Done or Watch methods should be used to
// ensure that a reply is available.
func (c *ReadCorrectableReply) Get() (*State, int, error) {
	c.Lock()
	defer c.Unlock()
	return c.State, c.level, c.err
}

// Done returns a channel that's closed when the correctable ReadCorrectable
// quorum call is done. A call is considered done when the quorum function has
// signaled that a quorum of replies was received or that the call returned an
// error.
func (c *ReadCorrectableReply) Done() <-chan struct{} {
	return c.donech
}

// Watch returns a channel that's closed when a reply or error at or above the
// specified level is available. If the call is done, the channel is closed
// disregardless of the specified level.
func (c *ReadCorrectableReply) Watch(level int) <-chan struct{} {
	ch := make(chan struct{})
	c.Lock()
	if level < c.level {
		close(ch)
		c.Unlock()
		return ch
	}
	c.watchers = append(c.watchers, &struct {
		level int
		ch    chan struct{}
	}{level, ch})
	c.Unlock()
	return ch
}

func (c *ReadCorrectableReply) set(reply *State, level int, err error, done bool) {
	c.Lock()
	if c.done {
		c.Unlock()
		panic("set(...) called on a done correctable")
	}
	c.State, c.level, c.err, c.done = reply, level, err, done
	if done {
		close(c.donech)
		for _, watcher := range c.watchers {
			if watcher != nil {
				close(watcher.ch)
			}
		}
		c.Unlock()
		return
	}
	for i := range c.watchers {
		if c.watchers[i] != nil && c.watchers[i].level <= level {
			close(c.watchers[i].ch)
			c.watchers[i] = nil
		}
	}
	c.Unlock()
}

/* Methods on Manager for correctable method ReadCorrectable */

type readCorrectableReply struct {
	nid   uint32
	reply *State
	err   error
}

func (m *Manager) readCorrectable(ctx context.Context, c *Configuration, corr *ReadCorrectableReply, args *ReadRequest) {
	replyChan := make(chan readCorrectableReply, c.n)

	for _, n := range c.nodes {
		go callGRPCReadCorrectable(ctx, n, args, replyChan)
	}

	var (
		replyValues = make([]*State, 0, c.n)
		clevel      = LevelNotSet
		reply       *State
		rlevel      int
		errCount    int
		quorum      bool
	)

	for {
		select {
		case r := <-replyChan:
			corr.NodeIDs = append(corr.NodeIDs, r.nid)
			if r.err != nil {
				errCount++
				break
			}
			replyValues = append(replyValues, r.reply)
			reply, rlevel, quorum = c.qspec.ReadCorrectableQF(replyValues)
			if quorum {
				corr.set(reply, rlevel, nil, true)
				return
			}
			if rlevel > clevel {
				clevel = rlevel
				corr.set(reply, rlevel, nil, false)
			}
		case <-ctx.Done():
			corr.set(reply, clevel, QuorumCallError{ctx.Err().Error(), errCount, len(replyValues)}, true)
			return
		}

		if errCount+len(replyValues) == c.n {
			corr.set(reply, clevel, QuorumCallError{"incomplete call", errCount, len(replyValues)}, true)
			return
		}
	}
}

func callGRPCReadCorrectable(ctx context.Context, node *Node, arg *ReadRequest, replyChan chan<- readCorrectableReply) {
	reply := new(State)
	start := time.Now()
	err := grpc.Invoke(
		ctx,
		"/dev.Register/ReadCorrectable",
		arg,
		reply,
		node.conn,
	)
	switch grpc.Code(err) { // nil -> codes.OK
	case codes.OK, codes.Canceled:
		node.setLatency(time.Since(start))
	default:
		node.setLastErr(err)
	}
	replyChan <- readCorrectableReply{node.id, reply, err}
}
