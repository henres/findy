import { Component, Prop, Vue } from 'vue-property-decorator';
import { Ressource } from '@/models/Ressource';
// import RessourceDetail from './RessourceDetail';
import { RessourceService } from '@/services/RessourceService'
import WithRender from '@/templates/ressources/list.ressources.html';

@WithRender
@Component
export default class RessourceList extends Vue {
    private ressources: Ressource[] = this.getRessources()
    private selectedRessource: Ressource | null = null
    private addingRessource: boolean = false;

    constructor() {
        super()
    }

    private created(): void {
        return
    }

    private deleteRessource(): void {
        return
    }

    private enableAddMode(): void {
        return
    }

    private getRessources(): Ressource[] {
        this.ressources = [];
        this.selectedRessource = null;
        return RessourceService.getRessources()
            .then((response: any) => {
                this.ressources = response.data
            }
        );
    }

    private ressourceChanged(arg: { ressource: Ressource, mode: string }): void {
        return
    }

    private onSelect(ressource: Ressource): Ressource[] {
        return []
    }

    private unSelect(): void {
        return
    }
}
