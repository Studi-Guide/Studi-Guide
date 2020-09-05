import {ElementRef, Renderer2} from '@angular/core';
import * as Hammer from 'hammerjs';

export class CanvasTouchHelper {

    public static currentZoom:
        { x: number; y: number; z: number };

    private static lastZoom :{
        x: number;
        y: number;
        z: number;
    };

    public static transformInOriginCoordinate(coordinates: { x: number; y: number}, canvas: HTMLCanvasElement) {
        console.log('Touch Interaction on: x' + coordinates.x + '... y:' + coordinates.y);
        const rect = canvas.getBoundingClientRect();
        const x = (coordinates.x - rect.left)/this.currentZoom.z;
        const y = (coordinates.y - rect.top)/this.currentZoom.z;

        console.log('Transformed to: x' + x + '... y:' + y);
        return {x, y};
    }

    public static RegisterPinch(renderer: Renderer2, canvasElement: ElementRef){
        const hammerTime = new Hammer(canvasElement.nativeElement, {});
        hammerTime.get('pinch').set({ enable: true });
        hammerTime.get('pan').set({ threshold: 0 });
        let fixHammerjsDeltaIssue = {
            x: 0,
            y: 0
        }
        let lastEvent;

        this.currentZoom = {
            x: 0,
            y: 0,
            z: 1,
        }

        this.lastZoom = {
            x: 0,
            y:0,
            z:1
        }
        let pinchZoomOrigin;

        hammerTime.on('doubletap', (event) => {
            if (this.currentZoom.z !== 1) {
                console.log('doubletab entered');
                const canvasHTMLElement = event.target as HTMLCanvasElement;

                canvasHTMLElement.style.transition = '0.3s';
                setTimeout(() => {
                    canvasHTMLElement.style.transition = 'none';
                }, 300)

                const originalSize = {
                    width: canvasHTMLElement.offsetWidth,
                    height: canvasHTMLElement.offsetHeight
                }

                const rect = (event.target as HTMLCanvasElement).getBoundingClientRect();
                const doubleTapOrigin = {x: event.center.x + rect.left, y: event.center.y + rect.top};
                console.log('doubleTapOrigin: x' + pinchZoomOrigin.x + '... y:' + pinchZoomOrigin.y);

                const d = this.scaleFrom(doubleTapOrigin, this.currentZoom.z, 1, originalSize)

                console.log('Scale back : x' + d.x + '... y:' + d.y +  '...z:' + d.z);
                this.currentZoom.x = 0;
                this.currentZoom.y = 0;
                this.currentZoom.z += d.z;

                this.lastZoom.x = this.currentZoom.x;
                this.lastZoom.y = this.currentZoom.y;
                this.lastZoom.z = this.currentZoom.z;

                this.update(this.currentZoom, canvasElement, renderer);
            }
        })

        hammerTime.on('pan', (event) => {
            console.log('pan entered');
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

            const transition = {
                x: event.deltaX - fixHammerjsDeltaIssue.x,
                y: event.deltaY - fixHammerjsDeltaIssue.y,
            };

            const canvasHTMLElement = event.target as HTMLCanvasElement;
            if (canvasHTMLElement != null && canvasHTMLElement !== undefined) {
                lastEvent = 'pan';
                    this.transistion(transition,
                        canvasElement, renderer, true)
                }
        })

        hammerTime.on('pinch', (event) => {

            const canvasHTMLElement = event.target as HTMLCanvasElement;
            const originalSize = {
                width: canvasHTMLElement.offsetWidth,
                height: canvasHTMLElement.offsetHeight
            }

            const d = this.scaleFrom(pinchZoomOrigin, this.lastZoom.z, this.lastZoom.z * event.scale, originalSize)
            this.currentZoom.x = d.x + this.lastZoom.x + event.deltaX;
            this.currentZoom.y = d.y + this.lastZoom.y + event.deltaY;
            this.currentZoom.z = d.z + this.lastZoom.z;
            lastEvent = 'pinch';
            this.update(this.currentZoom, canvasElement, renderer);
        })
        hammerTime.on('pinchstart', (event) =>  {
            const rect = (event.target as HTMLCanvasElement).getBoundingClientRect();
            pinchZoomOrigin = {x: event.center.x + rect.left, y: event.center.y + rect.top};
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

    public static transistion(
        transition: { x: number; y: number },
        canvasElement: ElementRef,
        renderer: Renderer2,
        isPan: boolean){

        const canvasHTMLElement = canvasElement.nativeElement as HTMLCanvasElement;
        const rect = canvasHTMLElement.getBoundingClientRect();
        const natizeElementSize = {
            width: rect.width,
            height: rect.height,
            top: rect.top,
            left: rect.left
        }

        const availableSize = {width: window.innerWidth, height: window.innerHeight};

        this.currentZoom.x = this.lastZoom.x + transition.x
        this.currentZoom.y = this.lastZoom.y + transition.y;
        console.log('Requested transition to x:' + this.currentZoom.x + '...y' + this.currentZoom.y)
        this.currentZoom = this.validateZoom(this.currentZoom, natizeElementSize, canvasHTMLElement, availableSize);
        this.update(this.currentZoom, canvasElement, renderer);
        if (!isPan) {
            this.lastZoom.x = this.currentZoom.x;
            this.lastZoom.y = this.currentZoom.y;
        }
    }

    private static validateZoom(currentZoom: { x: number; y: number; z: number;},
                                natizeElementSize: { width: number; height: number; top:number; left: number },
                                canvasElement: HTMLCanvasElement,
                                visibleSize: {width: number, height: number},) {

        const xTransitionMax = natizeElementSize.width - visibleSize.width;

        // allow a little bit of overdrive because of the tab and drawers
        const yTransitionMax = natizeElementSize.height - visibleSize.height * 0.8;

        const origin = this.transformInOriginCoordinate({x:0, y:0}, canvasElement)
        const x =  origin.x * (-1);
        const y =  origin.y * (-1);

        console.log('x..:' + x + 'y...' + y);
        const yFirstValue = y >= 0 ? Math.min(y, currentZoom.y) : Math.max(y, currentZoom.y);
        const xFirstValue = x >= 0 ? Math.min(x, currentZoom.x) : Math.max(x, currentZoom.x);

        console.log('min x..:' + x + 'max y...' + y);

        currentZoom.x = Math.max(xFirstValue, -xTransitionMax);
        currentZoom.y = Math.max(yFirstValue, -yTransitionMax);
        return currentZoom;
    }

    private static update(zoom: { x: number; y: number; z: number;},
                          element: ElementRef,
                          renderer: Renderer2) {
        console.log('Zoom to : x' + zoom.x + '... y:' + zoom.y +  '...z:' + zoom.z);

        renderer.setStyle(
            element.nativeElement,
            'transform',
            'translate3d(' + zoom.x + 'px, ' + zoom.y + 'px, 0) scale(' + zoom.z + ')');
    }

    private static scaleFrom(zoomOrigin, currentScale: number, newScale: number, originalSize) {
        const zoomDistance = newScale - currentScale
        const output = {
            x: 0, // * shift.x,
            y: 0, // * shift.y,
            z: zoomDistance
        }
        return output
    }
}
