export class CanvasResolutionConfigurator {

    public static setup(canvas: HTMLCanvasElement, width: number, height: number, scale:number, translationPos: TranslationPosition) {
        const dpr = this.GetPixelRatio();
        canvas.width = width * dpr;
        canvas.height = height * dpr;
        canvas.style.width = width + 'px';
        canvas.style.height = height + 'px';
        return this.getContext(canvas, scale, translationPos);
    }

    public static getContext(canvas: HTMLCanvasElement, scale: number, translationPos: TranslationPosition) {
        const dpr = this.GetPixelRatio();
        const ctx = canvas.getContext('2d');
        ctx.translate(translationPos.X, translationPos.Y);
        ctx.scale(dpr * scale, dpr * scale);
        return ctx;
    }

    public static GetPixelRatio() {
        return window.devicePixelRatio || 1;
    }
}

export class TranslationPosition {
    Y: number
    X: number
}
