using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Net;
using System.Net.Sockets;
using System.Security.Cryptography;
using System.Text;
using System.Threading;
using System.Windows.Forms;

namespace CryptoFrontendFileEther
{
    public partial class MainForm : Form
    {
        private string privateKey;
        public MainForm()
        {
            InitializeComponent();
        }
        private void logoutButton_Click(object sender, EventArgs e)
        {
            var loginForm = new LoginForm();
            loginForm.Show();
            this.Hide();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string fileHash = hashLabel.Text;
            StartClient();
        }

        private void MainForm_Load(object sender, EventArgs e)
        {

        }

        // The port number for the remote device.  
        private const int port = 8080;

        // ManualResetEvent instances signal completion.  
        private static ManualResetEvent connectDone =
            new ManualResetEvent(false);
        private static ManualResetEvent sendDone =
            new ManualResetEvent(false);
        private static ManualResetEvent receiveDone =
            new ManualResetEvent(false);

        // The response from the remote device.  
        private static String response = String.Empty;

        private static void StartClient()
        {
            // Connect to a remote device.  
            try
            {
                System.Net.IPAddress ipAddress = System.Net.IPAddress.Parse("192.168.43.173");
                IPEndPoint remoteEP = new IPEndPoint(ipAddress, port);

                // Create a TCP/IP socket.  
                Socket client = new Socket(ipAddress.AddressFamily,
                    SocketType.Stream, ProtocolType.Tcp);

                // Connect to the remote endpoint.  
                client.BeginConnect(remoteEP,
                    new AsyncCallback(ConnectCallback), client);
                connectDone.WaitOne();

                // Send test data to the remote device.  
                //Send(client, "This is a test<EOF>");
                //sendDone.WaitOne();

                // Receive the response from the remote device.  
                Receive(client);
                receiveDone.WaitOne();

                // Write the response to the console.  
                Console.WriteLine("Response received : {0}", response);

                // Release the socket.  
                client.Shutdown(SocketShutdown.Both);
                client.Close();

            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }

        private static void ConnectCallback(IAsyncResult ar)
        {
            try
            {
                // Retrieve the socket from the state object.  
                Socket client = (Socket)ar.AsyncState;

                // Complete the connection.  
                client.EndConnect(ar);

                Console.WriteLine("Socket connected to {0}",
                    client.RemoteEndPoint.ToString());

                // Signal that the connection has been made.  
                connectDone.Set();
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }

        private static void Receive(Socket client)
        {
            try
            {
                // Create the state object.  
                StateObject state = new StateObject();
                state.workSocket = client;

                // Begin receiving the data from the remote device.  
                client.BeginReceive(state.buffer, 0, StateObject.BufferSize, 0,
                    new AsyncCallback(ReceiveCallback), state);
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }

        private static void ReceiveCallback(IAsyncResult ar)
        {
            try
            {
                // Retrieve the state object and the client socket   
                // from the asynchronous state object.  
                StateObject state = (StateObject)ar.AsyncState;
                Socket client = state.workSocket;

                // Read data from the remote device.  
                int bytesRead = client.EndReceive(ar);

                if (bytesRead > 0)
                {
                    // There might be more data, so store the data received so far.  
                    state.sb.Write(state.buffer, 0, bytesRead);
                    state.responseStringBytes.Append(Encoding.ASCII.GetString(state.buffer, 0, bytesRead));

                    // Get the rest of the data.  
                    client.BeginReceive(state.buffer, 0, StateObject.BufferSize, 0,
                        new AsyncCallback(ReceiveCallback), state);
                }
                else
                {
                    // All the data has arrived; put it in response.  
                    if (state.sb.Length > 1)
                    {
                        MemoryStream result = new MainForm().FileDecrypt(state.sb);
                        response = bytesToString(result.GetBuffer());
                        Console.WriteLine(result);
                    }
                    // Signal that all bytes have been received.  
                    receiveDone.Set();
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }

        private static string bytesToString(byte[] input)
        {
            StringBuilder builder = new StringBuilder();

            builder.Append("[ ");

            for(int i = 0; i < input.Length; ++i)
            {
                builder.Append(input[i]);
                if(i != input.Length - 1)
                {
                    builder.Append(", ");
                }
            }

            builder.Append(" ]");

            return builder.ToString();
        }

        private static void Send(Socket client, String data)
        {
            // Convert the string data to byte data using ASCII encoding.  
            byte[] byteData = Encoding.ASCII.GetBytes(data);

            // Begin sending the data to the remote device.  
            client.BeginSend(byteData, 0, byteData.Length, 0,
                new AsyncCallback(SendCallback), client);
        }

        private static void SendCallback(IAsyncResult ar)
        {
            try
            {
                // Retrieve the socket from the state object.  
                Socket client = (Socket)ar.AsyncState;

                // Complete sending the data to the remote device.  
                int bytesSent = client.EndSend(ar);
                Console.WriteLine("Sent {0} bytes to server.", bytesSent);

                // Signal that all bytes have been sent.  
                sendDone.Set();
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }

        private byte[] generateRandomBytes()
        {
            return new byte[32];
        }

<<<<<<< HEAD
        private void test()
        {

        }

        private MemoryStream Encrypt()
        {
            byte[] toEncrypt = new byte[] { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16 };

            int keySize = 256;
            int blockSize = 128;

            byte[] passwordBytes = generateRandomBytes();
            byte[] ivBytes = new byte[blockSize / 8];

            Buffer.BlockCopy(passwordBytes, 0, ivBytes, 0, ivBytes.Length);
            MemoryStream fsOut = new MemoryStream();

            using (RijndaelManaged AES = new RijndaelManaged())
            {
                AES.BlockSize = blockSize;
                AES.KeySize = keySize;
                AES.Mode = CipherMode.CFB;
                AES.FeedbackSize = 8;
                AES.Padding = PaddingMode.None;

                AES.Key = passwordBytes;
                AES.IV = ivBytes;

                //var plainStream = new MemoryStream(toEncrypt, 0, toEncrypt.Length);

                // Create a decryptor to perform the stream transform.
                ICryptoTransform encryptor = AES.CreateEncryptor(AES.Key, AES.IV);

                // Create the streams used for encryption. 
                using (MemoryStream msEncrypt = new MemoryStream())
                {
                    using (CryptoStream csEncrypt = new CryptoStream(msEncrypt, encryptor, CryptoStreamMode.Write))
                    {
                        using (StreamWriter swEncrypt = new StreamWriter(csEncrypt))
                        {

                            //Write all data to the stream.
                            swEncrypt.Write(toEncrypt);
                        }

                        byte[] result = msEncrypt.ToArray();

                        fsOut.Write(result, 0, result.Length);
                    }
                }

            }
            
            return fsOut;
        }

=======
>>>>>>> 34afc086f42d41821d2532ca363b7f8335666da1
        private MemoryStream FileDecrypt(MemoryStream memStream)
        {
            int keySize = 256;
            int blockSize = 128;

            byte[] passwordBytes = generateRandomBytes();
            byte[] ivBytes = new byte[blockSize / 8];

            Buffer.BlockCopy(passwordBytes, 0, ivBytes, 0, ivBytes.Length);
            MemoryStream fsOut = new MemoryStream();

            memStream.Seek(0, SeekOrigin.Begin);
<<<<<<< HEAD
            fsOut.Seek(0, SeekOrigin.Begin);
=======
>>>>>>> 34afc086f42d41821d2532ca363b7f8335666da1

            using (RijndaelManaged AES = new RijndaelManaged())
            {
                AES.BlockSize = blockSize;
                AES.KeySize = keySize;
                AES.Mode = CipherMode.CFB;
                AES.FeedbackSize = 8;
                AES.Padding = PaddingMode.None;

                AES.Key = passwordBytes;
                AES.IV = ivBytes;

                using (var decryptor = AES.CreateDecryptor(AES.Key, AES.IV))
                using (var cs = new CryptoStream(memStream, decryptor, CryptoStreamMode.Read))
                {
                    cs.CopyTo(fsOut);
                }
            }

            /*
            RijndaelManaged AES = new RijndaelManaged();
            AES.KeySize = keySize;
            AES.BlockSize = blockSize;
            AES.Key = passwordBytes;
            AES.IV = ivBytes;
            AES.Padding = PaddingMode.None;
            AES.Mode = CipherMode.CFB;
            AES.FeedbackSize = 8;

            CryptoStream cs = new CryptoStream(memStream, AES.CreateDecryptor(), CryptoStreamMode.Read);
            MemoryStream fsOut = new MemoryStream();

            try
            {
                cs.CopyTo(fsOut);
            }
            catch (CryptographicException ex_CryptographicException)
            {
                Console.WriteLine("CryptographicException error: " + ex_CryptographicException.Message);
            }
            catch (Exception ex)
            {
                Console.WriteLine("Error: " + ex.Message);
            }
            
            try
            {
                cs.Close();
            }
            catch (Exception ex)
            {
                Console.WriteLine("Error by closing streams: " + ex.Message);
            }
            */

            return fsOut;
        }

        private void hashLabel_Click(object sender, EventArgs e)
        {

        }
    }
}
