import {ILocation, IRenderer} from '../../building-objects-if';


export class LocationRendererCanvas implements IRenderer {

    constructor(private location:ILocation) {
    }

    public get Location() : ILocation {
        return this.location;
    }

    render(renderingContext: CanvasRenderingContext2D) {
        renderingContext.font = '12px Arial';
        renderingContext.textAlign = 'center';
        renderingContext.fillStyle = '#000';
        renderingContext.fillText(this.location.Name, this.location.PathNode.Coordinate.X, this.location.PathNode.Coordinate.Y);
    }

    startAnimation(renderingContext: CanvasRenderingContext2D) {
    }

    stopAnimation(renderingContext: CanvasRenderingContext2D) {
    }

}