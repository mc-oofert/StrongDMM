package tool

import (
	"sdmm/dmapi/dm"
	"sdmm/dmapi/dmmap"
	"sdmm/dmapi/dmmap/dmmdata/dmmprefab"
	"sdmm/util"
)

type Modify interface {
	Dmm() *dmmap.Dmm
	UpdateCanvasByCoord(coord util.Point)
	SelectedPrefab() (*dmmprefab.Prefab, bool)
	CommitChanges(string)
}

// Add tool can be used to add prefabs to the map.
// During mouse moving when the tool is active a selected prefab will be added on every tile under the mouse.
// You can't add the same prefab twice on the same tile during the one OnStart -> OnStop cycle.
//
// Default: obj placed on top, area and turfs replaced.
// Alternative: obj replaced, area and turfs placed on top.
type Add struct {
	modify  Modify
	visuals Visuals

	// Objects will be replaced, turfs and areas will be added on top.
	altBehaviour bool

	tiles map[util.Point]bool
}

func NewAdd(modify Modify, visuals Visuals) *Add {
	return &Add{
		modify:  modify,
		visuals: visuals,
		tiles:   make(map[util.Point]bool),
	}
}

func (a *Add) OnStart(coord util.Point) {
	a.altBehaviour = isControlDown()
	a.OnMove(coord)
}

func (a *Add) OnMove(coord util.Point) {
	if prefab, ok := a.modify.SelectedPrefab(); ok && !a.tiles[coord] {
		a.tiles[coord] = true

		tile := a.modify.Dmm().GetTile(coord)

		if !a.altBehaviour {
			if dm.IsPath(prefab.Path(), "/area") {
				tile.InstancesRemoveByPath("/area")
			} else if dm.IsPath(prefab.Path(), "/turf") {
				tile.InstancesRemoveByPath("/turf")
			}
		} else if dm.IsPath(prefab.Path(), "/obj") {
			tile.InstancesRemoveByPath("/obj")
		}

		tile.InstancesAdd(prefab)
		tile.InstancesRegenerate()

		a.modify.UpdateCanvasByCoord(coord)
		a.visuals.MarkModifiedTile(coord)
	}
}

func (a *Add) OnStop(_ util.Point) {
	a.altBehaviour = false
	if len(a.tiles) != 0 {
		a.tiles = make(map[util.Point]bool, len(a.tiles))
		a.visuals.ClearModifiedTiles()
		go a.modify.CommitChanges("Add Atoms")
	}
}
