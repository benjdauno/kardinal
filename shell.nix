{pkgs}: let
  mergeShells = shells:
    pkgs.mkShell {
      shellHook = builtins.concatStringsSep "\n" (map (s: s.shellHook or "") shells);
      buildInputs = builtins.concatLists (map (s: s.buildInputs or []) shells);
      nativeBuildInputs = builtins.concatLists (map (s: s.nativeBuildInputs or []) shells);
      paths = builtins.concatLists (map (s: s.paths or []) shells);
    };
  manager_shell = pkgs.callPackage ./kardinal-manager/shell.nix {inherit pkgs;};
  cli_shell = pkgs.callPackage ./kardinal-cli/shell.nix {inherit pkgs;};
  cli_kontrol_api_shell = pkgs.callPackage ./libs/cli-kontrol-api/shell.nix {inherit pkgs;};
  kardinal_shell = with pkgs;
    pkgs.mkShell {
      buildInputs = [kubectl kustomize kubernetes-helm minikube istioctl tilt reflex];
      shellHook = ''
        source <(kubectl completion bash)
        source <(minikube completion bash)
        printf '\u001b[31m

                                          :::::
                                           :::::::
                                           ::   :::
                                          :::     ::
                                          ::   ::- :::
                                        :::         :::
                                       ::: :::    :::
                                     :::    ::    ::
                                   :::      ::   :::
                                 :::       :::   ::
                               :::        ::     ::
                            ::::       ::::     ::
                          ::::      ::::      :::
                       ::::::::::::::       ::::
                                       ::::::
                   :::::::::::::::::::::
               ::::::
            :::::
          :::



        \u001b[0m
        Starting Kardinal dev shell.
        \e[32m
        \e[0m
        '
      '';
    };
in
  mergeShells [manager_shell cli_shell kardinal_shell cli_kontrol_api_shell]
