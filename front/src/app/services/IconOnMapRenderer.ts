export class IconOnMapRenderer {

    private mapCanvasCtx: CanvasRenderingContext2D;
    private readonly icon: string;

    constructor(mapCanvasCtx:CanvasRenderingContext2D, icon:string) {
        this.mapCanvasCtx = mapCanvasCtx;
        this.icon = icon;
    }

    public render(x:number, y:number, width:number, height:number) {
        const image = new Image();
        image.onload = () => {
            this.mapCanvasCtx.drawImage(image, x, y, width, height);
        };
        image.src = 'assets/'+this.icon;
    }
}