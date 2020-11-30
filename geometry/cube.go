package geometry

func CubeVolume(n int) int {
  if n != 0 {
    return n * n * n, nil
  } else {
    return 0, erros.New("Zero lengh egde is not allowed")
  }
}