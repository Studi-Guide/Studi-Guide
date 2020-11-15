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

        const rect = canvas.getBoundingClientRect();
        const x = (coordinates.x - rect.left)/this.currentZoom.z;
        const y = (coordinates.y - rect.top)/this.currentZoom.z;


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


                const d = this.scaleFrom(doubleTapOrigin, this.currentZoom.z, 1, originalSize)


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
            if (lastEvent !== 'pan') {
                fixHammerjsDeltaIssue = {
                    x: event.deltaX,
                    y: event.deltaY
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
            if (event.scale === Infinity){
                return;
            }

            const canvasHTMLElement = event.target as HTMLCanvasElement;
            const originalSize = {
                width: canvasHTMLElement.offsetWidth,
                height: canvasHTMLElement.offsetHeight
            }

            const d = this.scaleFrom(pinchZoomOrigin, this.lastZoom.z, this.lastZoom.z * event.scale, originalSize)
            this.currentZoom.x = d.x + this.lastZoom.x + event.deltaX;
            this.currentZoom.y = d.y + this.lastZoom.y + event.deltaY;
            this.currentZoom.z = this.limitZoom(d.z, canvasElement);
            lastEvent = 'pinch';
            this.update(this.currentZoom, canvasElement, renderer);
        })
        hammerTime.on('pinchstart', (event) =>  {
            const rect = (event.target as HTMLCanvasElement).getBoundingClientRect();
            pinchZoomOrigin = {x: event.center.x + rect.left, y: event.center.y + rect.top};
            lastEvent = 'pinchstart';
        })

        hammerTime.on('panend', (event) => {
            this.lastZoom.x = this.currentZoom.x;
            this.lastZoom.y = this.currentZoom.y;
            lastEvent = 'panend';
        })

        hammerTime.on('pinchend', (event) => {
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
        this.currentZoom = this.validateZoom(this.currentZoom, canvasHTMLElement, availableSize, transition, !isPan);
        this.update(this.currentZoom, canvasElement, renderer);
        if (!isPan) {
            this.lastZoom.x = this.currentZoom.x;
            this.lastZoom.y = this.currentZoom.y;
        }
    }

    private static validateZoom(currentZoom: { x: number; y: number; z: number;},
                                canvasElement: HTMLCanvasElement,
                                visibleSize: {width: number, height: number},
                                transition: {x:number, y:number}, validateToMaxOnOver: boolean) {
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
        const yTransitionMax = (natizeElementSize.height * 1/currentZoom.z - visibleSize.height * 0.7)  * (-1) - 9000;

        // Introduce origin (9000/9000)
        const x =  9000 - ((origin.x + 25)* (Math.pow(currentZoom.z, 3)));
        const y =  9000 - ((origin.y + 25) *(Math.pow(currentZoom.z, 3)));

        const yTansistionMaxNegativ = -y;
        const xTansistionMaxNegativ = -x;

        let yzoomValue = currentZoom.y - 9000;
        if (yzoomValue < yTransitionMax && transition.y < 0) {

            yzoomValue = validateToMaxOnOver ? yTransitionMax : yvalueOld;
        }

        if (yzoomValue > yTansistionMaxNegativ && transition.y > 0) {
            const valueToSet = yzoomValue -(transition.y * 3/4);
            yzoomValue = validateToMaxOnOver ? yTansistionMaxNegativ : yvalueOld;
        }

        let xzoomValue = currentZoom.x- 9000;
        if (xzoomValue > xTansistionMaxNegativ && transition.x > 0) {
            const valueToSet=  xzoomValue -(transition.x* 3/4);
                xzoomValue = validateToMaxOnOver ? xTansistionMaxNegativ : xvalueOld;
        }

        if (xzoomValue < xTransitionMax && transition.x < 0) {
            xzoomValue  = validateToMaxOnOver ?  xTransitionMax: xvalueOld;
        }

        // console.log('Result Zoom x..:' + xzoomValue + ' y...' + yzoomValue);

        currentZoom.x = xzoomValue + 9000;
        currentZoom.y = yzoomValue + 9000;
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

    private static limitZoom(inOut:number, element: ElementRef) {
        const newSize = {x: element.nativeElement.getBoundingClientRect().width*(this.lastZoom.z+inOut),
            y: element.nativeElement.getBoundingClientRect().height*(this.lastZoom.z+inOut)};

        if (newSize.x < window.innerWidth*0.7 && newSize.y < window.innerHeight*0.4) {
            console.log(this.lastZoom.z, inOut);
            if (window.innerWidth*0.7 < window.innerWidth*0.4) {
                return 0.7;
            } else {
                return 0.4;
            }
        }

        return (this.lastZoom.z+inOut);
    }

    public static Zoom(inOut:number, element: ElementRef, renderer: Renderer2) {

        // const newSize = {x: element.nativeElement.getBoundingClientRect().width*(this.currentZoom.z+inOut),
        //     y: element.nativeElement.getBoundingClientRect().height*(this.currentZoom.z+inOut)};
        //
        // if (newSize.x < window.innerWidth*0.8 && newSize.y < window.innerHeight*0.5) {
        //     console.log(newSize, {x: window.innerWidth*0.8, y: window.innerHeight*0.5});
        //     newSize.x = window.innerWidth*0.8;
        //     newSize.y = window.innerHeight*0.5;
        // }

        this.currentZoom.z = this.limitZoom(inOut, element);// inOut;
        this.lastZoom.z = this.currentZoom.z;
        this.update(this.currentZoom, element, renderer);
    }
}
