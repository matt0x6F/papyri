import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import './index.css'
import VueFeather from 'vue-feather'

const app = createApp(App)
app.component(VueFeather.name, VueFeather)
app.use(router)
app.mount('#app');
