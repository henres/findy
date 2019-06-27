import { Component, Vue } from 'vue-property-decorator';
import RessourceList from '@/components/RessourceList';
import WithRender from '@/templates/app.html';

@WithRender
@Component({
  components: {
    RessourceList,
  },
})
export default class App extends Vue {}
