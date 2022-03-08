import { useEffect } from 'react'
import { grpc } from '@improbable-eng/grpc-web'
import useStore from '@/helpers/store'
import { CommunitiesService } from 'proto/zion/v1/communities_pb_service'
import { GetNodeInfoRequest } from 'proto/zion/v1/nodeinfo_pb'
import { Community, CreateCommunityRequest } from 'proto/zion/v1/communities_pb'
import { NodeInfoService } from 'proto/zion/v1/nodeinfo_pb_service'
import { CollectionsWriteRequest } from 'proto/identityhub/v1/collections_pb'
import { CollectionsService } from 'proto/identityhub/v1/collections_pb_service'
import { Request } from 'proto/identityhub/v1/requests_pb'

const nodeInfoRequest = new GetNodeInfoRequest()
const collectionsWriteRequest = new CollectionsWriteRequest()
const hubRequest = new Request()
collectionsWriteRequest.setRequest(hubRequest)
// hubRequest.setRequestid
// collectionsWriteRequest.

// const createCommunityRequest = new CreateCommunityRequest()
// const community = new Community()
// community.setName('Hello World Community! !! ! !!')
// community.setDescription('AMAZING DESCRIPTION ! ! !! ')
// createCommunityRequest.setPayload(community)

interface NodeInfoObject {
  pubkey: string
  balance: number
}

export const RPCTest = () => {
  const balance = useStore((s) => s.balance)
  const pubkey = useStore((s) => s.pubkey)
  useEffect(() => {
    console.log('Testing RPC connection')

    grpc.unary(NodeInfoService.GetNodeInfo, {
      request: nodeInfoRequest,
      host: 'https://mouse.zion.fyi',
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

    // grpc.unary(CollectionsService.CollectionsWrite, {
    //   request: collectionsWriteRequest,
    //   host: 'http://localhost:9090',
    //   onEnd: (res) => {
    //     console.log('CollectionsWrite response:', res)
    //   },
    // })

    // grpc.unary(CommunitiesService.CreateCommunity, {
    //   request: createCommunityRequest,
    //   host: 'http://localhost:9090',
    //   onEnd: (res) => {
    //     console.log('CreateCommunity response:', res)
    //   },
    // })
  }, [])
  return (
    <>
      <div
        className='bg-black h-screen w-screen flex flex-col justify-center items-center'
        style={{
          background:
            'linear-gradient(0deg, rgba(0,0,0,1) 0%, rgba(6,10,49,1) 100%)',
        }}
      >
        <h1 className='text-center pt-4 text-4xl font-extrabold text-white sm:pt-5 sm:text-6xl lg:pt-6 xl:text-6xl'>
          <span className=''>Zion</span>
          <span className='pt-2 pb-3 pl-4 bg-clip-text text-transparent bg-gradient-to-r from-teal-200 to-cyan-400 sm:pb-5'>
            Relay
          </span>
        </h1>
        <p className='mt-8 text-base text-gray-300 sm:text-xl lg:text-lg xl:text-xl'>
          {balance} sats
        </p>
        <p className='mt-4 text-gray-300 text-sm'>{pubkey}</p>
      </div>
    </>
  )
}
