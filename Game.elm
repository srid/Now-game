import Graphics.Element (..)
import Text (..)
import WebSocket
import Signal

helper = WebSocket.connect "ws://localhost:8000/dummy"    

main = Signal.map plainText (helper (Signal.constant "start"))
