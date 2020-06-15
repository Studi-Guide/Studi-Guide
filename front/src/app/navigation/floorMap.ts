import {Door, MapItem, SvgLocationName} from '../building-objects-if';
import {CanvasResolutionConfigurator} from '../services/CanvasResolutionConfigurator';
import {IconOnMapRenderer} from '../services/IconOnMapRenderer';

// @Injectable({
//   providedIn: 'root'
// })
export class FloorMap {
    public readonly pin: IconOnMapRenderer;

    public objectsToBeVisualized: MapItem[];
    public locationNames: SvgLocationName[];

    constructor(objectsToBeVisualized: MapItem[]) {
        this.objectsToBeVisualized = objectsToBeVisualized;
        this.pin = new IconOnMapRenderer('pin-sharp.png');
    }

    private renderDoor(map: CanvasRenderingContext2D, door:Door, color:string) {
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

    private renderRoomFrom(map: CanvasRenderingContext2D, mapItem:MapItem) {
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
            map.fillText(location.name,location.x,location.y);
        }
    }

    public renderFloorMap(map: CanvasRenderingContext2D) {
        for (const mapItem of this.objectsToBeVisualized) {
            this.renderRoomFrom(map, mapItem);
        }
        this.renderAllDoorsOfFloor(map);
        this.renderLocationNames(map);
    }

    public renderStartPin(map: CanvasRenderingContext2D, x:number, y:number, width:number, height:number) {
        this.pin.render(map, x, y, width, height);
    }
}
