import {IMapItem, IRenderer} from '../../building-objects-if';

export class MapItemRendererCanvas implements IRenderer {

    constructor(private mapItem:IMapItem) {
    }

    private doAnim = false;
    private animationCallback: () => void;

    public get MapItem() : IMapItem {
        return this.mapItem;
    }

    render(renderingContext: CanvasRenderingContext2D) {
        renderingContext.beginPath();
        this.renderSections(renderingContext);
        renderingContext.strokeStyle = '#FFF';
        renderingContext.fillStyle = this.mapItem.Color;
        renderingContext.stroke();
        renderingContext.fill();
        renderingContext.closePath();
    }

    startAnimation(renderingContext: CanvasRenderingContext2D, args?:any) {
        this.animationCallback = () => {
            if (this.doAnim !== true) {
                return;
            }
            renderingContext.save();
            const date = new Date();

            renderingContext.beginPath();
            this.renderSections(renderingContext);
            renderingContext.lineWidth = 1;
            renderingContext.strokeStyle = (date.getSeconds() % 2 === 0) ? '#FFF' : '#000';
            renderingContext.fillStyle = (date.getSeconds() % 2 === 0) ? '#000' : this.mapItem.Color;
            renderingContext.stroke();
            renderingContext.fill();
            renderingContext.closePath();

            renderingContext.restore();

            window.requestAnimationFrame(this.animationCallback);

            // route needs to be re-rendered after animation
            if (args !== undefined && args.renderer !== undefined && args.floor !== undefined) {
                for (const r of args.renderer) {
                    r.render(renderingContext, args);
                }
            }
        }

        this.doAnim = true;
        window.requestAnimationFrame(this.animationCallback);
    }

    stopAnimation(renderingContext: CanvasRenderingContext2D) {
        this.doAnim = false;
        this.animationCallback = null;
    }

    private renderSections(renderingContext:CanvasRenderingContext2D) {
        renderingContext.moveTo(this.mapItem.Sections[0].Start.X, this.mapItem.Sections[0].Start.Y);
        for (let i = 1; i < this.mapItem.Sections.length; i++) {
            renderingContext.lineTo(this.mapItem.Sections[i].Start.X, this.mapItem.Sections[i].Start.Y);
        }
        renderingContext.lineTo(this.mapItem.Sections[0].Start.X, this.mapItem.Sections[0].Start.Y);
    }

}