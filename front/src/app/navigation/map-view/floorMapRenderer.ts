import {Door, MapItem, Location} from '../../building-objects-if';
import {CanvasResolutionConfigurator} from '../../services/CanvasResolutionConfigurator';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';

// @Injectable({
//   providedIn: 'root'
// })
export class FloorMapRenderer {

    private readonly map: CanvasRenderingContext2D;
    private readonly mapCanvas: HTMLCanvasElement;
    private readonly objectsToBeVisualized: MapItem[];
    private readonly locationNames: Location[];

    constructor(objectsToBeVisualized: MapItem[],
                locationNames: Location[],
                ctx:CanvasRenderingContext2D,
                mapCanvas:HTMLCanvasElement) {
        this.map = ctx;
        this.mapCanvas = mapCanvas;
        this.objectsToBeVisualized = objectsToBeVisualized;
        this.locationNames = locationNames;
    }

    private renderDoor(door:Door, color:string) {
        this.map.beginPath();
        this.map.moveTo(door.Section.Start.X,door.Section.Start.Y);
        this.map.lineTo(door.Section.End.X,door.Section.End.Y);
        this.map.strokeStyle = color;
        this.map.stroke();
        this.map.closePath();
    }

    private renderAllDoorsOfFloor() {
        for (const mapItem of this.objectsToBeVisualized) {
            // TODO remove check if null
            if (mapItem.Doors != null && mapItem.Doors.length >= 1) {
                for (const door of mapItem.Doors) {
                    this.renderDoor(door, mapItem.Color);
                }
            }
        }
    }

    private renderRoomFrom(mapItem:MapItem){
        if (mapItem.Sections != null && mapItem.Sections.length > 0) {
            this.map.beginPath();
            this.map.moveTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
            for (let i = 1; i < mapItem.Sections.length; i++) {
                this.map.lineTo(mapItem.Sections[i].Start.X, mapItem.Sections[i].Start.Y);
            }
            this.map.lineTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
            this.map.strokeStyle = '#FFF';
            this.map.fillStyle = mapItem.Color;
            this.map.stroke();
            this.map.fill();
            this.map.closePath();
        }
    }

    private calcMapWidthHeight() {
        let mapHeightNeeded = 0;
        let mapWidthNeeded = 0;
        for (const mapItem of this.objectsToBeVisualized) {
            if (mapItem.Sections != null) {
                for (const section of mapItem.Sections) {
                    if (section.End.X > mapWidthNeeded) {
                        mapWidthNeeded = section.End.X;
                    }
                    if (section.End.Y > mapHeightNeeded) {
                        mapHeightNeeded = section.End.Y;
                    }
                }
            }
        }
        this.mapCanvas.height = mapHeightNeeded;
        this.mapCanvas.width = mapWidthNeeded+2;
    }

    private renderLocationNames() {
        this.map.font = '12px Arial';
        this.map.textAlign = 'center';
        this.map.fillStyle = '#000';
        for (const location of this.locationNames) {
            this.map.fillText(location.Name,location.PathNode.Coordinate.X,location.PathNode.Coordinate.Y);
        }
    }

    public renderFloorMap() {
        this.calcMapWidthHeight();
        for (const mapItem of this.objectsToBeVisualized) {
            this.renderRoomFrom(mapItem);
        }
        this.renderAllDoorsOfFloor();
        this.renderLocationNames();
    }
}
