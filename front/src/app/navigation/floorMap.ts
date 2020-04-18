import { Injectable } from '@angular/core';
import {Room, Door, Section, svgPath, MapItem} from "../building-objects-if";

// @Injectable({
//   providedIn: 'root'
// })
export class FloorMap {

    public objectsToBeVisualized: MapItem[];

    public calculatedRoomPaths: svgPath[];
    public calculatedDoorLines: svgPath[];

    public svgWidth: number;
    public svgHeight: number;

    constructor(objectsToBeVisualized: MapItem[]) {
        this.objectsToBeVisualized = objectsToBeVisualized;
        this.calculatedRoomPaths = [];
        this.calculatedDoorLines = [];
        this.svgWidth = 0;
        this.svgHeight = 0;
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
        for (const item of this.objectsToBeVisualized) {
            let itemShapePath:svgPath = {
                d : FloorMap.buildRoomSvgPathFromSections(item.Sections),
                fill : item.Color
            };
            this.calculatedRoomPaths.push(itemShapePath);
            // TODO remove check if null
            if (item.Doors != null && item.Doors.length >= 1) {
                for (const door of item.Doors) {
                    let doorLine:svgPath = {
                        d : FloorMap.buildDoorSvgLineFromSection(door),
                        fill : itemShapePath.fill
                    };
                    this.calculatedDoorLines.push(doorLine);
                }
            }
            for (const section of item.Sections) {
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
}