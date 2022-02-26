import useStore from '@/helpers/store'

export default function Instructions() {
  const balance = useStore((s) => s.balance)
  const pubkey = useStore((s) => s.pubkey)
  return (
    <>
      <div
        className='rounded-xl absolute max-w-lg px-4 py-2 text-sm shadow-xl pointer-events-none select-none md:text-base top-8 left-1/2 text-white transform -translate-x-1/2'
        style={{
          maxWidth: 'calc(100% - 28px)',
          backgroundColor: 'rgba(0,0,0,0.5)',
        }}
      >
        <p className='m-4 text-center'>Zion Relay</p>
      </div>
      <div
        className='rounded-xl absolute max-w-lg px-4 py-2 text-sm shadow-xl pointer-events-none select-none md:text-base top-32 left-1/2 text-white transform -translate-x-1/2'
        style={{
          maxWidth: 'calc(100% - 28px)',
          backgroundColor: 'rgba(0,0,0,0.3)',
        }}
      >
        <p className='m-4 text-center'>Balance: {balance} sats</p>
        <p className='m-4 text-center'>{pubkey}</p>
      </div>
    </>
  )
}
