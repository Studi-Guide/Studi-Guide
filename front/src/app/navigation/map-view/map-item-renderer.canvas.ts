import {IMapItem, IRenderer} from '../../building-objects-if';

export class MapItemRendererCanvas implements IRenderer {

    constructor(private mapItem:IMapItem) {
    }

    render(renderingContext: CanvasRenderingContext2D) {
        renderingContext.beginPath();
        renderingContext.moveTo(this.mapItem.Sections[0].Start.X, this.mapItem.Sections[0].Start.Y);
        for (let i = 1; i < this.mapItem.Sections.length; i++) {
            renderingContext.lineTo(this.mapItem.Sections[i].Start.X,this. mapItem.Sections[i].Start.Y);
        }
        renderingContext.lineTo(this.mapItem.Sections[0].Start.X, this.mapItem.Sections[0].Start.Y);
        renderingContext.strokeStyle = '#FFF';
        renderingContext.fillStyle = this.mapItem.Color;
        renderingContext.stroke();
        renderingContext.fill();
        renderingContext.closePath();
    }

    startAnimation(renderingContext: CanvasRenderingContext2D) {
    }

    stopAnimation(renderingContext: CanvasRenderingContext2D) {
    }

}