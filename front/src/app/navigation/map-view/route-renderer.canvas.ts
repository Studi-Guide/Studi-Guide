import {IRenderer} from '../../building-objects-if';
import {IReceivedRoute} from '../../route-objects-if';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';


export class RouteRendererCanvas implements IRenderer {

    constructor(private route:IReceivedRoute) {
    }

    render(renderingContext: CanvasRenderingContext2D, args?: any) {

        const pin = new IconOnMapRenderer('pin-sharp.png');
        if (this.route.Start.Floor === args.floor) {
            const x = this.route.Start.Node.Coordinate.X;
            const y = this.route.Start.Node.Coordinate.Y;
            pin.render(renderingContext,x-15, y-30, 30, 30);
        }
        const flag = new IconOnMapRenderer('flag-sharp.png');
        if (this.route.End.Floor === args.floor) {
            const x = this.route.End.Node.Coordinate.X;
            const y = this.route.End.Node.Coordinate.Y;
            flag.render(renderingContext, x-5, y-30, 30, 30);
        }


        const routeSections = this.route.RouteSections.filter(section => section.Floor === args.floor);
        if (routeSections != null) {
            for (const routeSection of routeSections) {
                renderingContext.strokeStyle = '#A00';
                renderingContext.lineWidth = 3;
                renderingContext.beginPath();
                renderingContext.moveTo(routeSection.Route[0].Coordinate.X,routeSection.Route[0].Coordinate.Y);
                for (let i = 1; i < routeSection.Route.length; i++) {
                    renderingContext.lineTo(routeSection.Route[i].Coordinate.X,routeSection.Route[i].Coordinate.Y);
                }
                renderingContext.stroke();
                renderingContext.closePath();
            }
        }
    }

    startAnimation(renderingContext: CanvasRenderingContext2D) {
    }

    stopAnimation(renderingContext: CanvasRenderingContext2D) {
    }

};