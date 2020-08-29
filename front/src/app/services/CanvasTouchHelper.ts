import {ElementRef, Renderer2} from '@angular/core';
import * as Hammer from 'hammerjs';

export class CanvasTouchHelper {

    private static currentZoom:
        { x: number; width: number; y: number; z: number;height: number };

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
        let fixHammerjsDeltaIssue;
        const pinchStart = { x: undefined, y: undefined }
        let lastEvent;

        const originalSize = {
            width: canvasElement.nativeElement.offsetWidth,
            height: canvasElement.nativeElement.offsetHeight
        }

        this.currentZoom = {
            x: 0,
            y: 0,
            z: 1,
            width: originalSize.width * 1,
            height: originalSize.height * 1,
        }

        const last = {
            x: this.currentZoom.x,
            y: this.currentZoom.y,
            z: this.currentZoom.z
        }

        hammerTime.on('doubletap', (event) => {
            canvasElement.nativeElement.style.transition = '0.3s';
            setTimeout(() => {
                canvasElement.nativeElement.style.transition = 'none';
            }, 300)

            const zoomOrigin = this.CalculateXY(
                {x: event.deltaX, y: event.deltaY},
                event.currentTarget as HTMLElement ?? event.target as HTMLElement);
            const d = this.scaleFrom(zoomOrigin, this.currentZoom.z, 1, originalSize)
            this.currentZoom.x += d.x;
            this.currentZoom.y += d.y;
            this.currentZoom.z += d.z;

            last.x = this.currentZoom.x;
            last.y = this.currentZoom.y;
            last.z = this.currentZoom.z;

            this.update(originalSize, this.currentZoom, canvasElement, renderer);
        })

        hammerTime.on('pan', (event) => {
            if (lastEvent !== 'pan') {
                fixHammerjsDeltaIssue = {
                    x: event.deltaX,
                    y: event.deltaY
                }
            }

            this.currentZoom.x = last.x + event.deltaX - fixHammerjsDeltaIssue.x;
            this.currentZoom.y = last.y + event.deltaY - fixHammerjsDeltaIssue.y;
            lastEvent = 'pan';
            this.update(originalSize, this.currentZoom, canvasElement, renderer);
        })

        hammerTime.on('pinch', (event) => {
            const d = this.scaleFrom(pinchZoomOrigin, last.z, last.z * event.scale, originalSize)
            this.currentZoom.x = d.x + last.x + event.deltaX;
            this.currentZoom.y = d.y + last.y + event.deltaY;
            this.currentZoom.z = d.z + last.z;
            lastEvent = 'pinch';
            this.update(originalSize, this.currentZoom, canvasElement, renderer);
        })

        let pinchZoomOrigin;
        hammerTime.on('pinchstart', (event) =>  {
            pinchStart.x = event.deltaX;
            pinchStart.y = event.deltaY;
            pinchZoomOrigin = this.getRelativePosition(event.currentTarget as HTMLElement ?? event.target as HTMLElement,
                { x: pinchStart.x, y: pinchStart.y }, originalSize, this.currentZoom.z);
            lastEvent = 'pinchstart';
        })

        hammerTime.on('panend', (event: MSGestureEvent) => {
            last.x = this.currentZoom.x;
            last.y = this.currentZoom.y;
            lastEvent = 'panend';
        })

        hammerTime.on('pinchend', (event: MSGestureEvent) => {
            last.x = this.currentZoom.x;
            last.y = this.currentZoom.y;
            last.z = this.currentZoom.z;
            lastEvent = 'pinchend';
        })
    }

    private static update(originalSize: { width: any; height: any },
                          zoom: { x: number; width: number; y: number; z: number; height: number },
                          element: ElementRef,
                          renderer: Renderer2) {
        this.currentZoom.height = originalSize.height * + zoom.z;
        this.currentZoom.width = originalSize.width * zoom.z;
        renderer.setStyle(element.nativeElement, 'zoom', zoom.z);
        element.nativeElement.scrollTo(zoom.x,zoom.y);
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

    private  static getCoords(elem: HTMLElement) { // crossbrowser version
        const box = elem.getBoundingClientRect();

        const body = document.body;
        const docEl = document.documentElement;

        const scrollTop = window.pageYOffset || docEl.scrollTop || body.scrollTop;
        const scrollLeft = window.pageXOffset || docEl.scrollLeft || body.scrollLeft;

        const clientTop = docEl.clientTop || body.clientTop || 0;
        const clientLeft = docEl.clientLeft || body.clientLeft || 0;

        const top  = box.top +  scrollTop - clientTop;
        const left = box.left + scrollLeft - clientLeft;

        return { x: Math.round(left), y: Math.round(top) };
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
