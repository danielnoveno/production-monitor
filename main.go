package main

import (
    "fmt"
    "os/exec"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func runPythonScript(script string) (string, error) {
    out, err := exec.Command("python", script).CombinedOutput()
    return string(out), err
}

func main() {
    a := app.New()
    w := a.NewWindow("Production Monitor - Dummy Simulation")
    w.Resize(fyne.NewSize(500, 300))

    productInput := widget.NewEntry()
    productInput.SetPlaceHolder("Scan/Enter Product Code")
    secondInput := widget.NewEntry()
    secondInput.SetPlaceHolder("Scan/Enter Batch/Worker Code")

    logLabel := widget.NewMultiLineEntry()
    logLabel.SetPlaceHolder("Log akan muncul di sini...")

    scanBtn := widget.NewButton("Simulate Python Scan", func() {
        code, err := runPythonScript("python_scripts/dummy_scan.py")
        if err != nil {
            logLabel.SetText(fmt.Sprintf("Error: %v", err))
        } else {
            productInput.SetText(code)
            logLabel.SetText(logLabel.Text + "\nBarcode generated: " + code)
        }
    })

    saveBtn := widget.NewButton("Save Record", func() {
        p := productInput.Text
        s := secondInput.Text
        if p == "" || s == "" {
            logLabel.SetText(logLabel.Text + "\n[ERROR] Product atau Second code kosong")
        } else {
            logLabel.SetText(logLabel.Text + fmt.Sprintf("\nSaved record: %s | %s", p, s))
            productInput.SetText("")
            secondInput.SetText("")
        }
    })

    w.SetContent(container.NewVBox(
        widget.NewLabel("Product Code:"),
        productInput,
        widget.NewLabel("Second Code:"),
        secondInput,
        scanBtn,
        saveBtn,
        logLabel,
    ))

    w.ShowAndRun()
}
