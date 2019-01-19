import java.io.*;
import java.net.Socket;
import java.util.*;

public class TcpClient {

    public final static int SOCKET_PORT = 8080;      // you may change this
    public final static String SERVER = "192.168.43.173";  // localhost
    public final static String FILE_TO_RECEIVED = "d:/curs04.pdf";  // you may change this, I give a
    // different name because i don't want to
    // overwrite the one used by server...

    public final static int FILE_SIZE = 6022386; // file size temporary hard coded
    // should bigger than the file to be downloaded

    private byte[] decryptChunk(byte[] chunkWithIV, byte[] key, int ivSize) {

        byte[] iv = Arrays.copyOfRange(chunkWithIV, chunkWithIV.length-ivSize, chunkWithIV.length);
        byte[] chunk = Arrays.copyOfRange(chunkWithIV, 0, chunkWithIV.length-ivSize);

        System.out.println(prettyPrintByteArray(chunk));
        System.out.println(prettyPrintByteArray(iv));

        return TcpClientDecrypt.decrypt(key, iv, chunk);
        //return TcpClientDecrypt.decrypt(key, new byte[]{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, chunkWithIV);
    }

    private byte[] requestPDF(String server, int port) throws IOException {

        ByteArrayOutputStream output = new ByteArrayOutputStream();
        Socket sock = null;

        try {
            sock = new Socket(server, port);
            System.out.println("Connecting...");

//            OutputStream sockOutput = sock.getOutputStream();
//
//            byte[] b = new byte[32];
//            new Random().nextBytes(b);
//
//            sockOutput.write(b, 0, b.length);

            int ivSize = 16;
            int chunkSize = 64;

            byte[] key = { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 };

            // receive file
            byte [] chunk  = new byte [chunkSize + ivSize];

            InputStream is = sock.getInputStream();

            int bytesRead = 0;

            while(bytesRead > -1) {
                // Decrypt each chunk
                bytesRead = is.read(chunk, 0, chunk.length);

                if(bytesRead == -1)
                    break;

                byte[] decrypted = decryptChunk(chunk, key, ivSize);

                output.write(decrypted);
            }

        }
        finally {
            if (sock != null) sock.close();
            output.close();
        }

        return output.toByteArray();
    }

    private String prettyPrintByteArray(byte[] bytes) {
        StringJoiner joiner = new StringJoiner(", ");

        for(int i = 0; i < bytes.length; ++i) {
            if (bytes[i] < 0) {
                int nofin = bytes[i] + 255;
                joiner.add(Integer.toString(nofin));
            }
            else {
                joiner.add(Byte.toString(bytes[i]));
            }
        }

        return joiner.toString();
    }

    public static void main (String [] args ) throws IOException {
        byte[] result = new TcpClient().requestPDF(SERVER, SOCKET_PORT);

        System.out.println(new TcpClient().prettyPrintByteArray(result));
    }
}