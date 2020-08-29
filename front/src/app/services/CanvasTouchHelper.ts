import {ElementRef, Renderer2} from '@angular/core';
import * as Hammer from 'hammerjs';

export class CanvasTouchHelper {

    private static currentZoom: number;

    public static CalculateXY(event:MouseEvent, canvasElement:HTMLElement) {
        const rect = canvasElement.getBoundingClientRect();
        const x = event.clientX/this.currentZoom - rect.left;
        const y = event.clientY/this.currentZoom - rect.top;
        console.log('Recognized Interaction with zoom ' + this.currentZoom + 'on ... x: ' + x + ' y: ' + y );

        return [x, y];
    }

    public static RegisterPinch(renderer: Renderer2, canvasElement: ElementRef){
        const defaultZoom = 1;
        this.currentZoom = defaultZoom; // Default original zoom
        const minZoom  = 0.3;
        const maxZoom  = 1.5;
        const zoomVelocity = 0.03;

        const hammerTime = new Hammer.Manager(canvasElement.nativeElement, { touchAction: 'pan-x pan-y' });
        const pinch = new Hammer.Pinch();

        hammerTime.add(pinch);

        // Tap recognizer with minimal 2 taps
        hammerTime.add(new Hammer.Tap({ event: 'doubletap', taps: 2 }) );

        hammerTime.on('pinchin', (event: MSGestureEvent) => {
            let newZoom = this.currentZoom - zoomVelocity;
            newZoom = newZoom < minZoom ? minZoom : newZoom;
            renderer.setStyle(canvasElement.nativeElement, 'zoom', newZoom);
            this.currentZoom = newZoom;
        });

        hammerTime.on('pinchout', (event: MSGestureEvent) => {
            let newZoom = this.currentZoom + zoomVelocity;
            newZoom = newZoom > maxZoom ? maxZoom : newZoom;
            renderer.setStyle(canvasElement.nativeElement, 'zoom', newZoom);
            this.currentZoom = newZoom;
        });

        hammerTime.on('doubletap', (event: MSGestureEvent)  => {
            renderer.setStyle(canvasElement.nativeElement, 'zoom', defaultZoom);
            this.currentZoom = defaultZoom;
        });
    }
}
