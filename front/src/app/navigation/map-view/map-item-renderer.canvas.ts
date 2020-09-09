import {IMapItem, IRenderer} from '../../building-objects-if';

export class MapItemRendererCanvas implements IRenderer {

    constructor(private mapItem:IMapItem) {
    }

    private doAnim = false;
    private animationCallback: () => void;

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
        this.animationCallback = () => {
            if (this.doAnim !== true) {
                return;
            }
            renderingContext.save();
            const date = new Date();

            renderingContext.beginPath();
            renderingContext.moveTo(this.mapItem.Sections[0].Start.X, this.mapItem.Sections[0].Start.Y);
            for (let i = 1; i < this.mapItem.Sections.length; i++) {
                renderingContext.lineTo(this.mapItem.Sections[i].Start.X, this.mapItem.Sections[i].Start.Y);
            }
            renderingContext.lineTo(this.mapItem.Sections[0].Start.X, this.mapItem.Sections[0].Start.Y);
            renderingContext.lineWidth = 1;
            renderingContext.strokeStyle = (date.getSeconds() % 2 === 0) ? '#FFF' : '#000';
            renderingContext.stroke();
            renderingContext.closePath();

            renderingContext.restore();

            window.requestAnimationFrame(this.animationCallback);
        }
    }

    stopAnimation(renderingContext: CanvasRenderingContext2D) {
        this.doAnim = false;
        this.animationCallback = null;
    }

}