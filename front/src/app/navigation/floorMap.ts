import { Injectable } from '@angular/core';
import {Room, Door, Section, svgPath, RoomName} from "../building-objects-if";

@Injectable({
    providedIn: 'root'
})
export class FloorMap {

    public objectsToBeVisualized: Room[];

    public calculatedRoomPaths: svgPath[];
    public calculatedDoorLines: svgPath[];

    public svgWidth: number;
    public svgHeight: number;
    public allRoomNames: RoomName[];

    constructor(objectsToBeVisualized: Room[]) {
        this.objectsToBeVisualized = objectsToBeVisualized;
        this.calculatedRoomPaths = [];
        this.calculatedDoorLines = [];
        this.svgWidth = 0;
        this.svgHeight = 0;
        this.allRoomNames = [];
    }

    private static buildDoorSvgLineFromSection(door:Door) : string {
        let path:string = 'M' + door.Section.Start.X + ' ' + door.Section.Start.Y;
        path += ' L' + door.Section.End.X + ' ' + door.Section.End.Y;
        return path;
    }

    private static buildRoomSvgPathFromSections(roomSections:Section[]) : string {
        let path_d:string = 'M';
        for (const section of roomSections) {
            if (path_d !== 'M') {
                path_d += 'L';
            }
            path_d += section.Start.X + ' ' + section.Start.Y + ' ';
        }
        path_d += 'Z';
        return path_d;
    }

    public calculateSvgPathsAndSvgWidthHeight() {
        for (const room of this.objectsToBeVisualized) {
            let roomShapePath:svgPath = {
                d : FloorMap.buildRoomSvgPathFromSections(room.Sections),
                fill : room.Color
            };
            this.calculatedRoomPaths.push(roomShapePath);
            // TODO remove check if null
            if (room.Doors != null && room.Doors.length >= 1) {
                for (const door of room.Doors) {
                    let doorLine:svgPath = {
                        d : FloorMap.buildDoorSvgLineFromSection(door),
                        fill : roomShapePath.fill
                    };
                    this.calculatedDoorLines.push(doorLine);
                }
            }
            for (const section of room.Sections) {
                if (section.End.X > this.svgWidth) {
                    this.svgWidth = section.End.X;
                }
                if (section.End.Y > this.svgHeight) {
                    this.svgHeight = section.End.Y;
                }
            }
            // bottom navigation bar overlays svg
            this.svgHeight += 1;
            this.svgWidth += 0.15;
        }
    }

    public collectAllRoomNames() {
        this.allRoomNames = []; // why is this initialization necessary?
        for (const room of this.objectsToBeVisualized) {
            this.allRoomNames.push(
                {
                    name: room.Name,
                    x: room.Sections[0].Start.X+6,
                    y: room.Sections[0].Start.Y+16
                }
            );
        }
    }
}