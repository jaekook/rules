package redis

import (
	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/redisutils"
	"github.com/project-flogo/rules/rete/internal/types"
)

//Holds a tuple reference and related state

type reteHandleImpl struct {
	types.NwElemIdImpl
	tuple    model.Tuple
	tupleKey model.TupleKey
	key      string
	status   string
	//jtRefs   types.JtRefsService
}

func newReteHandleImpl(nw types.Network, tuple model.Tuple, key, status string) types.ReteHandle {
	h1 := reteHandleImpl{}
	h1.initHandleImpl(nw, tuple, key, status)
	return &h1
}

func (hdl *reteHandleImpl) SetTuple(tuple model.Tuple) {
	hdl.tuple = tuple
	if tuple != nil {
		hdl.tupleKey = tuple.GetKey()
	}
}

func (hdl *reteHandleImpl) initHandleImpl(nw types.Network, tuple model.Tuple, key, status string) {
	hdl.SetID(nw)
	hdl.SetTuple(tuple)
	hdl.tupleKey = tuple.GetKey()
	hdl.key = key
	hdl.status = status
}

func (hdl *reteHandleImpl) GetTuple() model.Tuple {
	return hdl.tuple
}

func (hdl *reteHandleImpl) GetTupleKey() model.TupleKey {
	return hdl.tupleKey
}

func (hdl *reteHandleImpl) SetStatus(status string) {
	if hdl.key == "" {
		return
	}
	redisutils.GetRedisHdl().HSetNX(hdl.key, "status", status)
}

func (hdl *reteHandleImpl) GetStatus() string {
	return hdl.status
}

func (hdl *reteHandleImpl) AddJoinTableRowRef(joinTableRowVar types.JoinTableRow, joinTableVar types.JoinTable) {
	hdl.Nw.GetJtRefService().AddEntry(hdl, joinTableVar.GetName(), joinTableRowVar.GetID())
}

//Used when a rule is deleted. See Network.RemoveRule
func (hdl *reteHandleImpl) RemoveJoinTable(jtName string) {
	hdl.Nw.GetJtRefService().RemoveEntry(hdl, jtName)
}

func (hdl *reteHandleImpl) GetRefTableIterator() types.JointableIterator {
	refTblIterator := hdl.Nw.GetJtRefService().GetTableIterator(hdl)
	return refTblIterator
}