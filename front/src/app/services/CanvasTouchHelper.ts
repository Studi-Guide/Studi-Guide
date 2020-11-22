
export class CanvasTouchHelper {
    public static transformInOriginCoordinate(
        coordinates: { x: number; y: number},
        currentZoom: number,
        canvas: HTMLElement) {
        const rect = canvas.getBoundingClientRect();
        const x = (coordinates.x - rect.left)/currentZoom;
        const y = (coordinates.y - rect.top)/currentZoom;
        console.log('Transform x:' + coordinates.x + ' y:' + coordinates.y + '  to coordinate: x:' + x  + ' y:' + y);
        return {x, y};
    }
}
