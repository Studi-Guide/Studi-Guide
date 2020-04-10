import { Injectable } from '@angular/core';
import {Room, Section, svgPath, RoomName} from "../building-objects-if";

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

    private static buildDoorSvgLineFromSection(doorSection:Section) : string {
        let path:string = 'M' + doorSection.Start.X + ' ' + doorSection.Start.Y;
        path += ' L' + doorSection.End.X + ' ' + doorSection.End.Y;
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
                d : FloorMap.buildRoomSvgPathFromSections(room.sections),
                fill : room.Color
            };
            this.calculatedRoomPaths.push(roomShapePath);
            if (room.doors.length >= 1) {
                for (const door of room.doors) {
                    let doorLine:svgPath = {
                        d : FloorMap.buildDoorSvgLineFromSection(door),
                        fill : roomShapePath.fill
                    };
                    this.calculatedDoorLines.push(doorLine);
                }
            }
            for (const section of room.sections) {
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
                    name: room.name,
                    x: room.sections[0].Start.X+6,
                    y: room.sections[0].Start.Y+16
                }
            );
        }
    }
}