<template>
	<div class="courseitem">
		<div>
			<va-box
	        title="发布课题"
	        theme="box-info"
	        :isBorder="true"
	        widgetType="">
		        <div slot="content">
		          <form class="form-horizontal">
		            <va-input
		              :isHorizontal="true"
		              title="题目"
		              placeholder="请输入题目"
		              type="text"
		              vaId="horizontal-email"
		              >

		            </va-input>

		            <va-input
		              :isHorizontal="true"
		              title="类型"
		              placeholder="请输入课题类型"
		              type="text"
		              vaId="horizontal-email"
		              >

		            </va-input>

					<va-input
		              :isHorizontal="true"
		              title="要求"
		              placeholder="请输入课题要求"
		              type="text"
		              vaId="horizontal-email"
		              >

		            </va-input>

		            <va-textarea-group
			            title="内容"
			            :isDisabled="false"
			            placeholder="请输入课题内容..."
			          >
			        </va-textarea-group>

		          </form>

		        </div>

		        <div slot="footer">
		          <va-button :isBlock="false" theme="info" align="right" name="发布"></va-button>

		        </div>

		    </va-box>
		</div>
	</div>
</template>
<script>
	import VABox from './VABox'
	import VACheckBox from '../components/VACheckBox'
	import VAInput from '../components/VAInput'
	import VAButton from '../components/VAButton'
	import VATextareaGroup from '../components/VATextareaGroup'
	export default{
		name: 'VACourseItem',
		data(){
			return{

			}
		},
		components: {
		    'va-box': VABox,
		    'va-checkbox': VACheckBox,
		    'va-input': VAInput,
		    'va-button': VAButton,
		    'va-textarea-group': VATextareaGroup
		}
	}
</script>