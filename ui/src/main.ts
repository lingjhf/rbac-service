import { createApp } from 'vue'
import { i18n } from './i18n'
import App from './App.vue'
import { router } from './router'
import { pinia } from './store'
import 'uno.css'
import './styles.css'

const app = createApp(App)

app.use(router).use(pinia).use(i18n).mount('#app')

app.config.errorHandler = (err, instance, info) => {
  console.log('handle', err, info)
}
