const slotMachineService=require('./slotMachineService');

module.exports={
    
    Query:{
        ...slotMachineService.Query
    },
    Mutation:{
        ...slotMachineService.Mutation
    }
}
