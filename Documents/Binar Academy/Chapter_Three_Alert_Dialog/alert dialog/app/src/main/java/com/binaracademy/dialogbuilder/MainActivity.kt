package com.binaracademy.dialogbuilder

import android.app.AlertDialog
import android.app.Dialog
import android.content.DialogInterface
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.view.LayoutInflater
import android.widget.Button
import com.google.android.material.snackbar.Snackbar

class MainActivity : AppCompatActivity() {

    private lateinit var btnSimple : Button
    private lateinit var btnAlert : Button
    private lateinit var btnCustom : Button
    private lateinit var btnFragment : Button


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        btnSimple = findViewById(R.id.btnSimpleDialog)
        btnAlert = findViewById(R.id.btnAlertDialog)
        btnCustom = findViewById(R.id.btnCustomDialog)
        btnFragment = findViewById(R.id.btnFragmentDialog)


        btnSimple.setOnClickListener {
          AlertDialog.Builder(this)
                .setTitle("Testing Dialog")
                .setMessage("Pilih pilihan anda")
                .setCancelable(true)
                .show()
        }
        btnAlert.setOnClickListener {
        AlertDialog.Builder(this)
                .setTitle("Testing Dialog")
                .setMessage("Pilih pilihan anda")
                .setCancelable(true)
                .setPositiveButton("Pilihan satu"){
                        dialog, _ -> Snackbar.make(this.findViewById(R.id.parent),"Chapter Three",Snackbar.LENGTH_LONG).show()
                }
                .setNegativeButton("Tutup"){
                        dialog,_ -> dialog.dismiss()
                }
                .show()
        }

        btnCustom.setOnClickListener {
            val view = LayoutInflater.from(this).inflate(R.layout.custome_dialog,null,false)
            val dialog = AlertDialog.Builder(this)
            dialog.setView(view)
            dialog.setCancelable(false)
            val dialogRead = dialog.create()
            val btnOk = view.findViewById<Button>(R.id.btnOk)
            val btnCancel = view.findViewById<Button>(R.id.btnCancel)

            btnOk.setOnClickListener {
                Snackbar.make(this.findViewById(R.id.parent),"Custom Dialog SnackBar",Snackbar.LENGTH_LONG).show()
            }

            btnCancel.setOnClickListener {
                dialogRead.dismiss()
            }

                dialogRead.show()
        }

        btnFragment.setOnClickListener {
            val fragmentDialog = FragmentDialog()
            fragmentDialog.isCancelable = false
            fragmentDialog.show(supportFragmentManager,"ea")
        }



    }

}