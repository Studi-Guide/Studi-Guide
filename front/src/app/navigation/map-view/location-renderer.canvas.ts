import {ILocation, IRenderer} from '../../building-objects-if';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';


export class LocationRendererCanvas implements IRenderer {

    constructor(private location: ILocation) {
    }

    public get Location(): ILocation {
        return this.location;
    }

    render(renderingContext: CanvasRenderingContext2D) {
        if (this.location.Icon.length > 0) {
            const r = new IconOnMapRenderer('svg/' + this.location.Icon + '.svg');
            r.render(renderingContext, this.location.PathNode.Coordinate.X, this.location.PathNode.Coordinate.Y, 20, 20);
        } else {
        renderingContext.font = '12px Arial';
        renderingContext.textAlign = 'center';
        renderingContext.fillStyle = '#000';
        renderingContext.fillText(this.location.Name, this.location.PathNode.Coordinate.X, this.location.PathNode.Coordinate.Y);
        }
    }

    startAnimation(renderingContext: CanvasRenderingContext2D) {
    }

    stopAnimation(renderingContext: CanvasRenderingContext2D) {
    }

}
