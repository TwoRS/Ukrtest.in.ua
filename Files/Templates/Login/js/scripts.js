<% define "vue.js" %>


var Forum = Vue.extend({
    template: '<p>This is Forum!</p>'
})

var Chat = Vue.extend({
    template: '<p>This is Chat!</p>'
})


var App = Vue.extend({
    el: '#app',
    data: function() {
        return {
            Myvalue1: 'Roman'
        }
    }
})

var router = new VueRouter({
    history: true
})

router.map({
    '/forum': {
        component: Forum
    },
    '/chat': {
        component: Chat
    }
})

router.start(App, '#app')
    //var App = Vue.extend({})



<% end %>