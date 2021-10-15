import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'element-plus/dist/index.css'
import {ElLoading} from 'element-plus';

createApp(App)
    .use(store)
    .use(router)
    .use(ElLoading)
    .mount('#app')
