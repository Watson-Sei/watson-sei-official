import Vue from 'vue';
import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import enUS from '@kangc/v-md-editor/lib/lang/en-US';
VueMarkdownEditor.use(vuepressTheme);
VueMarkdownEditor.lang.use('en-US', enUS);
Vue.use(VueMarkdownEditor);
