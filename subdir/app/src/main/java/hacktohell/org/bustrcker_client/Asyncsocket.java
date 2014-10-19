package hacktohell.org.bustrcker_client;

/**
 * Created by gowtham on 10/19/2014.
 */
import android.app.Service;
import android.content.Intent;
import android.os.AsyncTask;
import android.os.IBinder;
import android.util.Log;

import com.koushikdutta.async.AsyncNetworkSocket;
import com.koushikdutta.async.AsyncServer;
import com.koushikdutta.async.AsyncServerSocket;
import com.koushikdutta.async.AsyncSocket;
import com.koushikdutta.async.ByteBufferList;
import com.koushikdutta.async.DataEmitter;
import com.koushikdutta.async.callback.CompletedCallback;
import com.koushikdutta.async.callback.DataCallback;
import com.koushikdutta.async.callback.ListenCallback;

import java.io.BufferedOutputStream;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.Socket;


public class Asyncsocket extends AsyncTask<Void,Void,Void>{
    private String url;
    private int port;
    private byte[] data;
    Socket s;

    Asyncsocket(String url1,int port1,byte[] data1){
        url=url1;
        port=port1;
        data=data1;

    }
    @Override
    protected Void doInBackground(Void... params) {

        try {
            s = new Socket(url, port);
            BufferedReader in = new BufferedReader(new InputStreamReader(s.getInputStream()));
            BufferedOutputStream out = new BufferedOutputStream(s.getOutputStream());
               if (s != null)
                if (s.isConnected()) {
                    out.write(data);
                    out.flush();
                }
        }
        catch(Exception e){

        }
        return null;

    }
}
