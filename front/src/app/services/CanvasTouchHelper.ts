import {ElementRef, Renderer2} from '@angular/core';
import * as Hammer from 'hammerjs';

export class CanvasTouchHelper {

    public static currentZoom:
        { x: number; width: number; y: number; z: number;height: number };

    private static lastZoom :{
        x: number;
        y: number;
        z: number;
    };
    private static originalSize: { width: number; height: number };

    public static CalculateXY(coordinates: { x: number; y: number }, canvasElement: HTMLElement) {
        const rect = canvasElement.getBoundingClientRect();
        const x = coordinates.x/this.currentZoom.z - rect.left;
        const y = coordinates.y/this.currentZoom.z - rect.top;
        console.log('Recognized Interaction with zoom ' + this.currentZoom.z + 'on ... x: ' + x + ' y: ' + y );

        return [x, y];
    }

    public static RegisterPinch(renderer: Renderer2, canvasElement: ElementRef){
        const hammerTime = new Hammer(canvasElement.nativeElement, {});
        hammerTime.get('pinch').set({ enable: true });
        hammerTime.get('pan').set({ threshold: 0 });
        let fixHammerjsDeltaIssue = {
            x: 0,
            y: 0
        }
        const pinchStart = { x: undefined, y: undefined }
        let lastEvent;

        this.originalSize = {
            width: canvasElement.nativeElement.offsetWidth,
            height: canvasElement.nativeElement.offsetHeight
        }

        this.currentZoom = {
            x: 0,
            y: 0,
            z: 1,
            width: this.originalSize.width * 1,
            height: this.originalSize.height * 1,
        }

        this.lastZoom = {
            x: 0,
            y:0,
            z:1
        }

        hammerTime.on('doubletap', (event) => {
            if (this.currentZoom.z !== 1) {
                canvasElement.nativeElement.style.transition = '0.3s';
                setTimeout(() => {
                    canvasElement.nativeElement.style.transition = 'none';
                }, 300)

                const zoomOrigin = this.getRelativePosition(
                    event.currentTarget as HTMLElement ?? event.target as HTMLElement,
                    {x: event.center.x, y: event.center.y}, this.originalSize, this.currentZoom.z);


                // const zoomOrigin = this.CalculateXY(
                //   {x: event.deltaX, y: event.deltaY},
                //    event.currentTarget as HTMLElement ?? event.target as HTMLElement);
                const d = this.scaleFrom(zoomOrigin, this.currentZoom.z, 1, this.originalSize)
                this.currentZoom.x += d.x;
                this.currentZoom.y += d.y;
                this.currentZoom.z += d.z;

                this.lastZoom.x = this.currentZoom.x;
                this.lastZoom.y = this.currentZoom.y;
                this.lastZoom.z = this.currentZoom.z;

                this.update(this.originalSize, this.currentZoom, canvasElement, renderer);
            }
        })

        hammerTime.on('pan', (event) => {
            if (lastEvent !== 'pan') {
                fixHammerjsDeltaIssue = {
                    x: event.deltaX,
                    y: event.deltaY
                }
            }
            else {
                fixHammerjsDeltaIssue = {
                    x: 0,
                    y: 0
                }
            }

            const xTransition = event.deltaX - fixHammerjsDeltaIssue.x;
            const yTransition = event.deltaY - fixHammerjsDeltaIssue.y;
            this.transistion(xTransition, yTransition, canvasElement, renderer)
            lastEvent = 'pan';
        })

        hammerTime.on('pinch', (event) => {
            const d = this.scaleFrom(pinchZoomOrigin, this.lastZoom.z, this.lastZoom.z * event.scale, this.originalSize)
            this.currentZoom.x = d.x + this.lastZoom.x + event.deltaX;
            this.currentZoom.y = d.y + this.lastZoom.y + event.deltaY;
            this.currentZoom.z = d.z + this.lastZoom.z;
            lastEvent = 'pinch';
            this.update(this.originalSize, this.currentZoom, canvasElement, renderer);
        })

        let pinchZoomOrigin;
        hammerTime.on('pinchstart', (event) =>  {
            pinchStart.x = event.deltaX;
            pinchStart.y = event.deltaY;
            pinchZoomOrigin = this.getRelativePosition(event.currentTarget as HTMLElement ?? event.target as HTMLElement,
                { x: pinchStart.x, y: pinchStart.y }, this.originalSize, this.currentZoom.z);
            lastEvent = 'pinchstart';
        })

        hammerTime.on('panend', (event: MSGestureEvent) => {
            this.lastZoom.x = this.currentZoom.x;
            this.lastZoom.y = this.currentZoom.y;
            lastEvent = 'panend';
        })

        hammerTime.on('pinchend', (event: MSGestureEvent) => {
            this.lastZoom.x = this.currentZoom.x;
            this.lastZoom.y = this.currentZoom.y;
            this.lastZoom.z = this.currentZoom.z;
            lastEvent = 'pinchend';
        })
    }

    public static transistion(xCoordinate: number, yCoordinate: number, canvasElement: ElementRef, renderer: Renderer2){
        this.currentZoom.x = this.lastZoom.x + xCoordinate
        this.currentZoom.y = this.lastZoom.y + yCoordinate;
        this.update(this.originalSize, this.currentZoom, canvasElement, renderer);
    }

    private static update(originalSize: { width: any; height: any },
                          zoom: { x: number; width: number; y: number; z: number; height: number },
                          element: ElementRef,
                          renderer: Renderer2) {
        this.currentZoom.height = originalSize.height * + zoom.z;
        this.currentZoom.width = originalSize.width * zoom.z;
        renderer.setStyle(
            element.nativeElement,
            'transform',
            'translate3d(' + this.currentZoom.x + 'px, ' + this.currentZoom.y + 'px, 0) scale(' + this.currentZoom.z + ')');
        // renderer.setStyle(element.nativeElement, 'zoom', zoom.z);
        // element.nativeElement.scrollTo(zoom.x,zoom.y);
        // renderer.setStyle(element.nativeElement, )
        // element.style.transform =
        //    ;
    }

    private static getRelativePosition(element: HTMLElement, point, originalSize, scale) {
        const domCoords = this.getCoords(element);

        const elementX = point.x - domCoords.x;
        const elementY = point.y - domCoords.y;

        const relativeX = elementX / (originalSize.width * scale / 2) - 1;
        const relativeY = elementY / (originalSize.height * scale / 2) - 1;
        return { x: relativeX, y: relativeY }
    }

    private static getCoords(elem: HTMLElement) { // crossbrowser version
        const box = elem.getBoundingClientRect();
        return { x: Math.round(box.left), y: Math.round(box.top) };
    }

    private static scaleFrom(zoomOrigin, currentScale: number, newScale: number, originalSize) {
        const currentShift = this.getCoordinateShiftDueToScale(originalSize, currentScale);
        const newShift = this.getCoordinateShiftDueToScale(originalSize, newScale)

        const zoomDistance = newScale - currentScale

        const shift = {
            x: currentShift.x - newShift.x,
            y: currentShift.y - newShift.y,
        }

        const output = {
            x: zoomOrigin.x * shift.x,
            y: zoomOrigin.y * shift.y,
            z: zoomDistance
        }
        return output
    }

    private static getCoordinateShiftDueToScale(size, scale: number){
        const newWidth = scale * size.width;
        const newHeight = scale * size.height;
        const dx = (newWidth - size.width) / 2
        const dy = (newHeight - size.height) / 2
        return {
            x: dx,
            y: dy
        }
    }
}
