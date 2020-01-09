<template>
	<div class="usercourse">
		<div>
			<va-todo-list></va-todo-list>
		</div>
	</div>
</template>
<script>
	import VATodoList from '../widgets/VATodoList.vue'
	export default{
		name: 'VAUserCourse',
		data(){
			return{

			}
		},
		components: {
		    'va-todo-list': VATodoList
		}
	}
</script>