import {ICampus} from '../building-objects-if';
export class CampusViewModel {

    private readonly _campus: ICampus
    private readonly _addressString: string;
    constructor(campus: ICampus) {
        this._campus = campus;
        this._addressString = this.getAddressString(this._campus)
    }

    public get Address() {
        return this._addressString;
    }

    public get Name() {
        return this._campus.Name;
    }

    public get ShortName() {
        return this._campus.ShortName;
    }

    private getAddressString(campus: ICampus) {
        let addressString = '';
        if (campus.edges.Address !== null) {
            const addressGroup = groupBy(campus.edges.Address, i => i.Street);
            for (const street of Object.keys(addressGroup)) {
                const group = addressGroup[street]
                if (addressString !== '') {
                    addressString += '; '
                }

                let addressNumber = group[0].Number
                if (group.length > 1) {
                    addressNumber += '-' + group[group.length - 1].Number
                }

                addressString += street + ' ' + addressNumber + ', '
                addressString += addressGroup[street][0].PLZ + ' ' + addressGroup[street][0].City
            }
        }

        return addressString;
    }
}

const groupBy = <T, K extends keyof any>(list: T[], getKey: (item: T) => K) =>
    list.reduce((previous, currentItem) => {
        const group = getKey(currentItem);
        if (!previous[group]) previous[group] = [];
        previous[group].push(currentItem);
        return previous;
    }, {} as Record<K, T[]>);

