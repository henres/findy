import { Component, Prop, Vue, Watch, Emit } from 'vue-property-decorator'
import { Ressource } from '@/models/Ressource'
import WithRender from '@/templates/ressources/details.ressources.html';

@WithRender
@Component({})
export default class RessourceDetail extends Vue {
    @Prop() private ressource: Ressource | null = null
    private addingRessource = !this.ressource
    private editingRessource: Ressource | null = null

    @Watch('ressource') public onRessourceChange(value: string, oldValueL: string) {
        return
    }

    $refs!: {
        id: HTMLElement
        name: HTMLElement
    }

    addRessource() {
        return
    }

    @Emit('unselect') public clear() {
        this.editingRessource = null;
    }

    private cloneIt() {
        return
    }

    private created() {
        return
    }

    @Emit('ressourceChanged') public emitRefresh(mode: string, ressource: Ressource) {
        this.clear();
    }

    private mounted() {
        return
    }
}
