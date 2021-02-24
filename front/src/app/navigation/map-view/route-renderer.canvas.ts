import {IRenderer} from '../../building-objects-if';
import {IReceivedRoute} from '../../route-objects-if';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';


export class RouteRendererCanvas implements IRenderer {

    constructor(private route: IReceivedRoute) {
    }

    public get Route(): IReceivedRoute {
        return this.route;
    }

    render(renderingContext: CanvasRenderingContext2D, args?: any) {
        const pin = new IconOnMapRenderer('assets/pin-sharp.png');
        if (this.route.Start.Floor === args.floor) {
            const x = this.route.Start.Node.Coordinate.X;
            const y = this.route.Start.Node.Coordinate.Y;
            pin.render(renderingContext, x - 15, y - 30, 30, 30);
        }

        const flag =  new IconOnMapRenderer( 'assets/pin-red.png');
        if (this.route.End.Floor === args.floor) {
            const x = this.route.End.Node.Coordinate.X;
            const y = this.route.End.Node.Coordinate.Y;
            flag.render(renderingContext, x - 30,  y - 40, 60, 60);
        }


        const routeSections = this.route.RouteSections.filter(section => section.Floor === args.floor);
        if (routeSections != null) {
            // draw route line
            for (const routeSection of routeSections) {
                renderingContext.strokeStyle = '#A00';
                renderingContext.lineWidth = 3;
                renderingContext.beginPath();
                renderingContext.moveTo(routeSection.Route[0].Coordinate.X, routeSection.Route[0].Coordinate.Y);
                for (let i = 1; i < routeSection.Route.length; i++) {
                    renderingContext.lineTo(routeSection.Route[i].Coordinate.X, routeSection.Route[i].Coordinate.Y);
                }
                renderingContext.stroke();
                renderingContext.closePath();
            }

            // draw route length
            for (const routeSection of routeSections) {
                const numberOfPathNodes: number = routeSection.Route.length;
                const value: number = routeSection.Distance;
                const x: number = routeSection.Route[Math.round((numberOfPathNodes - 1) / 2)].Coordinate.X;
                const y: number = routeSection.Route[Math.round((numberOfPathNodes - 1) / 2)].Coordinate.Y;
                const font = '14px Arial';
                const width = renderingContext.measureText(value.toString()).width + 14;
                const height = parseInt('14px Arial', 10) + 14;
                renderingContext.fillStyle = '#A00';
                renderingContext.fillRect(x - width / 2, y - height / 2, width, 20);
                renderingContext.font = font;
                renderingContext.fillStyle = '#FFF';
                renderingContext.fillText(value.toString(), x, y);
            }

        }
    }

    startAnimation(renderingContext: CanvasRenderingContext2D) {
    }

    stopAnimation(renderingContext: CanvasRenderingContext2D) {
    }

}
