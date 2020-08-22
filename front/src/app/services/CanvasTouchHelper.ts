export class CanvasTouchHelper {
    public static CalculateXY(event:MouseEvent, canvasElement:HTMLElement) {
        const rect = canvasElement.getBoundingClientRect();
        const x = event.clientX - rect.left;
        const y = event.clientY - rect.top;
        console.log('Recognized Interaction on ... x: ' + x + ' y: ' + y);
        return [x, y];
    }
}
