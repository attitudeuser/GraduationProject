import Vue from 'vue'
import Router from 'vue-router'
import Hello from 'components/Hello.vue'
import Sample from 'components/Sample.vue'
import DashboardV1 from 'examples/Dashboard.v1.vue'
import DashboardV2 from 'examples/Dashboard.v2.vue'
import InfoBoxExample from 'examples/InfoBoxExample'
import ChartExample from 'examples/ChartExample'
import AlertExample from 'examples/AlertExample'
import ModalExample from 'examples/ModalExample'
import WidgetsExample from 'examples/WidgetsExample'
import APIExample from 'examples/APIExample'
import VALogin from 'widgets/VALogin.vue'
import VAUserInfor from 'widgets/VAUserInfor.vue'
import VAUserCourse from 'widgets/VAUserCourse.vue'
import VACourseItem from 'widgets/VACourseItem.vue'
import VACommunication from 'widgets/VACommunication.vue'

// UI Element Groups
import General from 'pages/ui-elements/General.vue'
import Icons from 'pages/ui-elements/Icons.vue'
import Buttons from 'pages/ui-elements/Buttons.vue'
import Sliders from 'pages/ui-elements/Sliders.vue'
import Timeline from 'pages/ui-elements/Timeline.vue'
import Modals from 'pages/ui-elements/Modals.vue'

// forms
import GeneralElements from 'pages/forms/GeneralElements.vue'
import AdvancedElements from 'pages/forms/AdvancedElements.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [{
      path: '/',
      name: 'home',
      redirect: '/user/profile',
    },
    {
      path: '/auth',
      name: 'auth',
      component: () => import('../views/UserAuth.vue')
    }, 
    {
      path: '/reset',
      name: 'reset',
      component: () => import('../views/Reset.vue')
    }, 
    {
      path: '/user',
      name: 'user',
      redirect: 'profile',
      component: () => import('../views/DashboardLayout.vue'),
      children: [{
          path: 'profile',
          name: 'profile',
          component: () => import('../views/Profile.vue')
        },
        {
          path: 'course',
          name: 'course',
          component: () => import('../views/Course.vue')
        },
        {
          path: 'publish',
          name: 'publish',
          component: () => import('../views/Publish.vue')
        },
        {
          path: 'communication',
          name: 'communication',
          component: () => import('../views/Communication.vue')
        }
      ]
    },
    {
      path: '/VAUserInfor',
      name: 'VAUserInfor',
      component: VAUserInfor
    },
    {
      path: '/VAUserCourse',
      name: 'VAUserCourse',
      component: VAUserCourse
    },
    {
      path: '/VACourseItem',
      name: 'VACourseItem',
      component: VACourseItem
    },
    {
      path: '/VACommunication',
      name: 'VACommunication',
      component: VACommunication
    },
    {
      path: '/sample',
      name: 'Sample',
      component: Sample
    },
    {
      path: '/dashboard/v1',
      name: 'DashboardV1',
      component: DashboardV1
    },
    {
      path: '/dashboard/v2',
      name: 'DashboardV2',
      component: DashboardV2
    },
    {
      path: '/examples/infobox',
      name: 'InfoBoxExample',
      component: InfoBoxExample
    },
    {
      path: '/examples/chart',
      name: 'ChartExample',
      component: ChartExample
    },
    {
      path: '/examples/alert',
      name: 'AlertExample',
      component: AlertExample
    },
    {
      path: '/examples/modal',
      name: 'ModalExample',
      component: ModalExample
    },
    {
      path: '/examples/widgets',
      name: 'WidgetsExample',
      component: WidgetsExample
    },
    {
      path: '/examples/api-example',
      name: 'APIExample',
      component: APIExample
    },
    {
      path: '/ui-elements/general',
      name: 'General',
      component: General
    },
    {
      path: '/ui-elements/icons',
      name: 'Icons',
      component: Icons
    },
    {
      path: '/ui-elements/buttons',
      name: 'Buttons',
      component: Buttons
    },
    {
      path: '/ui-elements/sliders',
      name: 'Sliders',
      component: Sliders
    },
    {
      path: '/ui-elements/timeline',
      name: 'Timeline',
      component: Timeline
    },
    {
      path: '/ui-elements/modals',
      name: 'Modals',
      component: Modals
    },
    {
      path: '/forms/general-elements',
      name: 'GeneralElements',
      component: GeneralElements
    },
    {
      path: '/forms/advanced-elements',
      name: 'AdvancedElements',
      component: AdvancedElements
    }
  ],
  linkActiveClass: 'active'
})
