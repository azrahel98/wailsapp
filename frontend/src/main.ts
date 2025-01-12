import {createApp} from 'vue'
import App from './App.vue'
import {router} from './router/router';


import './style.css';



const app = createApp(App);


  app.use(router);


  app.config.errorHandler = (err, vm, info) => {
  
    console.error(`[Global Error Handler]: ${info}`, err);
  };
  

  app.mount('#app');
