export class CanvasResolutionConfigurator {

    public static setup(canvas: HTMLCanvasElement, width: number, height: number) {
        const dpr = this.GetPixelRatio();
        canvas.width = width * dpr;
        canvas.height = height * dpr;
        canvas.style.width = width + 'px';
        canvas.style.height = height + 'px';
        const ctx = canvas.getContext('2d');
        ctx.scale(dpr, dpr);
        return ctx;
    }

    public static GetPixelRatio() {
        return window.devicePixelRatio || 1;
    }
}
