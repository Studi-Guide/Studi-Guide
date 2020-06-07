import {Door, MapItem, SvgLocationName} from '../building-objects-if';
import {CanvasResolutionConfigurator} from '../services/CanvasResolutionConfigurator';
import {IconOnMapRenderer} from '../services/IconOnMapRenderer';

// @Injectable({
//   providedIn: 'root'
// })
export class FloorMap {

    private readonly mapCanvas: HTMLCanvasElement;
    private map: CanvasRenderingContext2D;

    public readonly pin: IconOnMapRenderer;

    public objectsToBeVisualized: MapItem[];
    public locationNames: SvgLocationName[];

    constructor(objectsToBeVisualized: MapItem[]) {
        this.mapCanvas = document.getElementById('map') as HTMLCanvasElement;
        this.map = CanvasResolutionConfigurator.setup(this.mapCanvas);
        this.objectsToBeVisualized = objectsToBeVisualized;
        this.pin = new IconOnMapRenderer(this.map,'pin-sharp.png');
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
            this.map.fillText(location.name,location.x,location.y);
        }
    }

    private clearMapCanvas() {
        this.map.clearRect(0, 0, this.mapCanvas.width, this.mapCanvas.height);
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
