<template>
    <Transition name="slide-fade">
        <div v-if="show" class="reg-window" @click.self="closePage">
            <div class="create-task auth-form">
                <div class="space-betw" style="max-width: 96%;">
                    <div class="logo create" style="justify-content: flex-start;">
                        <span style="color:#434343">NeSU.</span><span style="color:#87C644">Еду</span><span style="color:#9ED5F0">*</span>
                    </div>
                    <img src="/images/check-mark.svg" style="max-height: 50px; margin-top: 0.7%; cursor: pointer;" @click="createTask">
                </div>

                <div style="display: flex; flex-direction: column; margin-left: 3%; max-width: 94%; gap: 5px;">
                    <input class="create-task-input" type="text" placeholder="Введите название" v-model="taskName">
                    <div class="space-betw" style="max-width: 55%; margin: 0;">
                        <input class="class-numb create-task-input" type="text" placeholder="аудитория" v-model="classNumber">
                        <input class="time create-task-input" type="text" placeholder="время" v-model="time">
                    </div>
                    
                    <div class="space-betw" style="max-width: 100%; margin: 0; align-items: center;">
                        <textarea class="desc create-task-input" type="text" placeholder="Описание (при необходимости)" v-model="desc"></textarea>
                        <div class="price">
                            <input class="create-task-input" style="width: 80%; margin-left: auto;" type="text" placeholder="? + 50 руб." v-model="price">
                        </div>
                        
                    </div> 
                </div>
            </div>
        </div>
    </Transition>
</template>

<script>
import { axios_utils } from '../../axios_utils'
export default {
    data: () => ({
        show: false,
        taskName: '',
        classNumber: '',
        time: '',
        price: '',
        desc: '',
    }),
    
    methods: {
        closePage: function() {
            this.show = false;
            this.taskName = '';
            this.classNumber = '';
            this.time = '';
            this.price = '';
            this.desc = '';
        },

        createTask: function() {
            var sessionID = sessionStorage.getItem("jwt");

            var hours = this.time.split(' ')
            var date = new Date();
            date.setHours(hours[0], hours[1], 0);

            var task = {
                title: this.taskName,
                description: this.desc,
                audience: this.classNumber,
                price: Number(this.price) + 50,
                deadline: date.toISOString(),
            }
            axios_utils.createTask(task, sessionID).then(result => {
                this.closePage();
                this.$emit('updateTasksList');
            });
        },
    },
}
</script>