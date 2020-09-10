import {IDoor, IMapItem, ILocation, IPathNode, ISection} from '../../building-objects-if';


export class FloorMapRenderer {
    public readonly objectsToBeVisualized: IMapItem[];
    public readonly locationNames: ILocation[];

    constructor(objectsToBeVisualized: IMapItem[],
                locationNames: ILocation[]) {
        this.objectsToBeVisualized = objectsToBeVisualized;
        this.locationNames = locationNames;
    }

    private renderDoor(map: CanvasRenderingContext2D, door:IDoor, color:string) {
        map.beginPath();
        map.moveTo(door.Section.Start.X,door.Section.Start.Y);
        map.lineTo(door.Section.End.X,door.Section.End.Y);
        map.strokeStyle = color;
        map.stroke();
        map.closePath();
    }

    private renderAllDoorsOfFloor(map: CanvasRenderingContext2D) {
        for (const mapItem of this.objectsToBeVisualized) {
            // TODO remove check if null
            if (mapItem.Doors != null && mapItem.Doors.length >= 1) {
                for (const door of mapItem.Doors) {
                    this.renderDoor(map, door, mapItem.Color);
                }
            }
        }
    }

    private renderRoomFrom(map: CanvasRenderingContext2D, mapItem:IMapItem) {
        if (mapItem.Sections != null && mapItem.Sections.length > 0) {
            map.beginPath();
            map.moveTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
            for (let i = 1; i < mapItem.Sections.length; i++) {
                map.lineTo(mapItem.Sections[i].Start.X, mapItem.Sections[i].Start.Y);
            }
            map.lineTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
            map.strokeStyle = '#FFF';
            map.fillStyle = mapItem.Color;
            map.stroke();
            map.fill();
            map.closePath();
        }
    }

    private renderLocationNames(map: CanvasRenderingContext2D) {
        map.font = '12px Arial';
        map.textAlign = 'center';
        map.fillStyle = '#000';
        for (const location of this.locationNames) {
            map.fillText(location.Name,location.PathNode.Coordinate.X,location.PathNode.Coordinate.Y);
        }
    }

    public renderFloorMap(map: CanvasRenderingContext2D) {
        for (const mapItem of this.objectsToBeVisualized) {
            this.renderRoomFrom(map, mapItem);
        }
        this.renderAllDoorsOfFloor(map);
        this.renderLocationNames(map);
    }
}
