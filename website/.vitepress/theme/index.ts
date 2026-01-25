import { h } from 'vue'
import Theme from 'vitepress/theme'
import './style.css'
import Layout from './Layout.vue'

import FeatureMatrix from '../components/FeatureMatrix.vue'
import HowItWorks from '../components/HowItWorks.vue'
import UseCases from '../components/UseCases.vue'
import CtaSection from '../components/CtaSection.vue'

export default {
  extends: Theme,
  Layout: Layout,
  enhanceApp({ app }) {
    app.component('FeatureMatrix', FeatureMatrix)
    app.component('HowItWorks', HowItWorks)
    app.component('UseCases', UseCases)
    app.component('CtaSection', CtaSection)
  }
}
