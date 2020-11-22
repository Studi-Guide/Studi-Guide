
export class CanvasTouchHelper {
    public static transformInOriginCoordinate(coordinates: { x: number; y: number}, currentZoom: number, canvas: HTMLCanvasElement) {

        const rect = canvas.getBoundingClientRect();
        const x = (coordinates.x - rect.left)/currentZoom;
        const y = (coordinates.y - rect.top)/currentZoom;
        return {x, y};
    }
}
