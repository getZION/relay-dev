import { useEffect } from 'react'
import { grpc } from '@improbable-eng/grpc-web'
import useStore from '@/helpers/store'
import { CommunitiesService } from 'proto/zion/v1/communities_pb_service'
import { GetNodeInfoRequest } from 'proto/zion/v1/nodeinfo_pb'
import { Community, CreateCommunityRequest } from 'proto/zion/v1/communities_pb'
import { NodeInfoService } from 'proto/zion/v1/nodeinfo_pb_service'

const nodeInfoRequest = new GetNodeInfoRequest()
const createCommunityRequest = new CreateCommunityRequest()
const community = new Community()
community.setName('Hello World Community! !! ! !!')
community.setDescription('AMAZING DESCRIPTION ! ! !! ')
createCommunityRequest.setPayload(community)

interface NodeInfoObject {
  pubkey: string
  balance: number
}

export const RPCTest = () => {
  useEffect(() => {
    console.log("Let's test RPC connection and community creation")

    grpc.unary(NodeInfoService.GetNodeInfo, {
      request: nodeInfoRequest,
      host: 'http://localhost:9090',
      onEnd: (res) => {
        if (res.message) {
          const info = res.message.toObject() as NodeInfoObject
          console.log(info.pubkey)
          console.log(info.balance)
          useStore.setState({
            pubkey: info.pubkey,
            balance: info.balance,
          })
        } else {
          console.log('error', res.statusMessage)
        }
      },
    })

    grpc.unary(CommunitiesService.CreateCommunity, {
      request: createCommunityRequest,
      host: 'http://localhost:9090',
      onEnd: (res) => {
        console.log('CreateCommunity response:', res)
      },
    })
  }, [])
  return <></>
}