//SPDX-License-Identifier:MIT
pragma solidity = 0.8.0;
contract Oracle{
    struct DataSourceNode {
        uint256 DataImportance;//数据重要性
        uint256 NodeFailureRate;//节点故障率
        uint256 Source;//来源渠道
        uint256 Integrity;//完整性
        uint256 Utilization;//使用率
        uint256 Result;//可信结果
        uint256 NumberOfEvaluated;//被评估次数

    }

    //使用mapping(uint256 => struct),根据节点Id存储对应的struct
    mapping(uint256 => DataSourceNode) public IdToNode;


    //增加节点
    function AddNode(uint256 NodeId) public {
        //创建实例       
        DataSourceNode memory newNode;
        
        IdToNode[NodeId] = newNode;

    }
    //评估节点评分
    function Grade(
        uint256 NodeId,
        uint256 _DataImportance,
        uint256 _NodeFailureRate,
        uint256 _Source,
        uint256 _Integrity,
        uint256 _Utilization,
        uint256 _Result
        )public {
            if(IdToNode[NodeId].NumberOfEvaluated == 0){
                //如果struct并没有传入过参数，则直接赋值
                IdToNode[NodeId].DataImportance = _DataImportance;
                IdToNode[NodeId].NodeFailureRate = _NodeFailureRate;
                IdToNode[NodeId].Source = _Source;
                IdToNode[NodeId].Integrity = _Integrity;
                IdToNode[NodeId].Utilization = _Utilization;
                IdToNode[NodeId].Result = _Result*100;
                IdToNode[NodeId].NumberOfEvaluated = 1;

            }else{
                //否则，计算公式为(分数*评估次数)+输入的参数                
                uint256 Number = IdToNode[NodeId].NumberOfEvaluated;
                IdToNode[NodeId].DataImportance = (IdToNode[NodeId].DataImportance*Number + _DataImportance)/(Number+1);
                IdToNode[NodeId].NodeFailureRate = (IdToNode[NodeId].NodeFailureRate*Number +  _NodeFailureRate)/(Number+1);
                IdToNode[NodeId].Source = (IdToNode[NodeId].Source*Number + _Source)/(Number+1);
                IdToNode[NodeId].Integrity = (IdToNode[NodeId].Integrity*Number + _Integrity)/(Number+1);
                IdToNode[NodeId].Utilization = (IdToNode[NodeId].Utilization*Number + _Utilization)/(Number+1);

                IdToNode[NodeId].Result = (IdToNode[NodeId].Result*Number + _Result*100)/(Number+1);
                //评估次数加1
                IdToNode[NodeId].NumberOfEvaluated += 1;
                
            }
        }
    //获取节点可信结果
    function GetResult(uint256 NodeId) public view returns(bool ){
        if(IdToNode[NodeId].Result>=51){
            return true;
        }else{
            return false;
        }

    }

    //删除节点信息
    function DeleteInformation(uint256 NodeId) public {
        DataSourceNode memory newNode;
        IdToNode[NodeId] = newNode;       
    }
}