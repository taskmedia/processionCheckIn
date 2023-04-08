import { createApp } from 'vue';
import App from './App.vue';

import router from './router';
import store from './store';

import BaseButton from './components/base/BaseButton.vue';
import BaseCard from './components/base/BaseCard.vue';
import TheHeader from './components/layout/TheHeader.vue';

const app = createApp(App);

// base components
app.component('base-button', BaseButton);
app.component('base-card', BaseCard);
app.component('the-header', TheHeader);

app.use(store);
app.use(router);
app.mount('#app');
