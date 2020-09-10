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

        const availableSize = {width: window.innerWidth, height: window.innerHeight};
        this.currentZoom = this.validateZoom(this.currentZoom, canvasHTMLElement, availableSize, transition);
        this.update(this.currentZoom, canvasElement, renderer);
        if (!isPan) {
            this.lastZoom.x = this.currentZoom.x;
            this.lastZoom.y = this.currentZoom.y;
        }
    }

    private static validateZoom(currentZoom: { x: number; y: number; z: number;},
                                canvasElement: HTMLCanvasElement,
                                visibleSize: {width: number, height: number},
                                transition: {x:number, y:number}) {
        const rect = canvasElement.getBoundingClientRect();
        const natizeElementSize = {
            width: rect.width,
            height: rect.height,
            top: rect.top,
            left: rect.left
        }

        const xvalueOld = currentZoom.x - 9000;
        const yvalueOld = currentZoom.y - 9000;
        currentZoom.x = this.lastZoom.x + transition.x
        currentZoom.y = this.lastZoom.y + transition.y;
        const origin = this.transformInOriginCoordinate({x:0, y:0}, canvasElement)
        const xTransitionMax = (natizeElementSize.width *1/currentZoom.z- visibleSize.width) * (-1) - 9000;

        // allow a little bit of overdrive because of the tab and drawers
        const yTransitionMax = (natizeElementSize.height * 1/currentZoom.z - visibleSize.height * 0.9)  * (-1) - 9000;

        // Introduce origin (9000/9000)
        const x =  9000 - ((origin.x + 25)* (Math.pow(currentZoom.z, 3)));
        const y =  9000 - ((origin.y + 25) *(Math.pow(currentZoom.z, 3)));

        const yTansistionMaxNegativ = -y;
        const xTansistionMaxNegativ = -x;

        let yzoomValue = currentZoom.y - 9000;
        if (yzoomValue < yTransitionMax && transition.y < 0) {
            console.log('Zoom Limit y_max reached. Reset value '   + yzoomValue + ' to: ' + yTransitionMax);
            yzoomValue = yvalueOld;
        }

        if (yzoomValue > yTansistionMaxNegativ && transition.y > 0) {
            const valueToSet = yzoomValue -(transition.y * 3/4);
            console.log('Zoom Limit y_min reached. Reset value '   + yzoomValue + ' to: ' + valueToSet);
            yzoomValue = yvalueOld;
        }

        let xzoomValue = currentZoom.x- 9000;
        if (xzoomValue > xTansistionMaxNegativ && transition.x > 0) {
            const valueToSet=  xzoomValue -(transition.x* 3/4);
            console.log('Zoom Limit x_min reached. Reset value '   + (xzoomValue) + ' to: ' + (valueToSet));
                xzoomValue = xvalueOld;
        }

        if (xzoomValue < xTransitionMax && transition.x < 0) {
            console.log('Zoom Limit x_max reached. Reset value '   + xzoomValue + ' to: ' + xTransitionMax);
            xzoomValue  = xvalueOld;
        }

        // console.log('Result Zoom x..:' + xzoomValue + ' y...' + yzoomValue);

        currentZoom.x = xzoomValue + 9000;
        currentZoom.y = yzoomValue + 9000;
        console.log('Result Zoom x..:' + currentZoom.x + ' y...' + currentZoom.y);
        return currentZoom;
    }

    private static update(zoom: { x: number; y: number; z: number;},
                          element: ElementRef,
                          renderer: Renderer2) {
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
