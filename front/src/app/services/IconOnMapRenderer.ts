export class IconOnMapRenderer {

    private readonly icon: string;

    constructor(icon:string) {
        this.icon = icon;
    }

    public render(mapCanvasCtx:CanvasRenderingContext2D, x:number, y:number, width:number, height:number) {
        const image = new Image();
        image.onload = () => {
            mapCanvasCtx.drawImage(image, x, y, width, height);
        };
        image.src = 'assets/'+this.icon;
    }
}
