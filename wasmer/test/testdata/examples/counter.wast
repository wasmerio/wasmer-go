(module
  (type (;0;) (func (param i32)))
  (type (;1;) (func (param i32 i32)))
  (type (;2;) (func (param i32) (result i32)))
  (type (;3;) (func (param i32 i32 i32) (result i32)))
  (type (;4;) (func (param i32 i32) (result i32)))
  (type (;5;) (func))
  (type (;6;) (func (result i32)))
  (type (;7;) (func (param i32) (result i64)))
  (type (;8;) (func (param i32 i32 i32)))
  (type (;9;) (func (param i32 i32 i32 i32)))
  (type (;10;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;11;) (func (param i32 i32 i32 i32 i32)))
  (type (;12;) (func (param i32 i32 i32 i32 i32 i32 i32) (result i32)))
  (type (;13;) (func (param i64 i32 i32) (result i32)))
  (type (;14;) (func (param i32 i32 i32 i32 i32 i32) (result i32)))
  (import "env" "inc" (func $inc (type 0)))
  (import "env" "get" (func $get (type 6)))
  (import "env" "copy_from_reg" (func $copy_from_reg (type 0)))
  (import "env" "copy_to_reg" (func $copy_to_reg (type 0)))
  (func $__wasm_call_ctors (type 5))
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h27b62dc0d035ddf3E (type 7) (param i32) (result i64)
    i64.const 7549865886324542212)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h4139b7e5fe9883cbE (type 7) (param i32) (result i64)
    i64.const -941380383357200164)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h64cc3526f0a264d5E (type 7) (param i32) (result i64)
    i64.const -670765639137414048)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17he5dd53c4163a6f10E (type 7) (param i32) (result i64)
    i64.const 1229646359891580772)
  (func $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h3575d213ce7d211eE (type 4) (param i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    i32.load
    local.tee 0
    i32.load offset=8
    local.set 3
    local.get 0
    i32.load
    local.set 0
    local.get 2
    local.get 1
    call $_ZN4core3fmt9Formatter10debug_list17h2471317d64372e8dE
    block  ;; label = @1
      local.get 3
      i32.eqz
      br_if 0 (;@1;)
      loop  ;; label = @2
        local.get 2
        local.get 0
        i32.store offset=12
        local.get 2
        local.get 2
        i32.const 12
        i32.add
        i32.const 1048600
        call $_ZN4core3fmt8builders8DebugSet5entry17hdfa2936ae55a6785E
        drop
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        local.get 3
        i32.const -1
        i32.add
        local.tee 3
        br_if 0 (;@2;)
      end
    end
    local.get 2
    call $_ZN4core3fmt8builders9DebugList6finish17hdbd1049d6264d19cE
    local.set 0
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h37cfe269a54a2708E (type 4) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.set 0
    block  ;; label = @1
      local.get 1
      call $_ZN4core3fmt9Formatter15debug_lower_hex17hfff82ceb2b4262ffE
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      call $_ZN4core3fmt3num52_$LT$impl$u20$core..fmt..LowerHex$u20$for$u20$i8$GT$3fmt17hac36f7f0417b1c1dE
      return
    end
    block  ;; label = @1
      local.get 1
      call $_ZN4core3fmt9Formatter15debug_upper_hex17h0a841ca4788fd580E
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      call $_ZN4core3fmt3num52_$LT$impl$u20$core..fmt..UpperHex$u20$for$u20$i8$GT$3fmt17hb957d38375b60145E
      return
    end
    local.get 0
    local.get 1
    call $_ZN4core3fmt3num3imp51_$LT$impl$u20$core..fmt..Display$u20$for$u20$u8$GT$3fmt17ha887097ae9f6ec79E)
  (func $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17hbcb23e4f920baed2E (type 4) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.set 0
    block  ;; label = @1
      local.get 1
      call $_ZN4core3fmt9Formatter15debug_lower_hex17hfff82ceb2b4262ffE
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      call $_ZN4core3fmt3num53_$LT$impl$u20$core..fmt..LowerHex$u20$for$u20$i32$GT$3fmt17h92e35f33f4a021aeE
      return
    end
    block  ;; label = @1
      local.get 1
      call $_ZN4core3fmt9Formatter15debug_upper_hex17h0a841ca4788fd580E
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      call $_ZN4core3fmt3num53_$LT$impl$u20$core..fmt..UpperHex$u20$for$u20$i32$GT$3fmt17he210bef9ff0aa51fE
      return
    end
    local.get 0
    local.get 1
    call $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$u32$GT$3fmt17h4f3293445fab7cb7E)
  (func $_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17haf2d0b9687871f57E (type 4) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    local.get 1
    call $_ZN42_$LT$str$u20$as$u20$core..fmt..Display$GT$3fmt17hfdf78a4ae71e6a9bE)
  (func $_ZN4core6result13unwrap_failed17h1984afaa28d103e8E (type 5)
    (local i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 0
    global.set 0
    local.get 0
    i32.const 16
    i32.store offset=12
    local.get 0
    i32.const 1048616
    i32.store offset=8
    local.get 0
    i32.const 52
    i32.add
    i32.const 1
    i32.store
    local.get 0
    i32.const 36
    i32.add
    i32.const 2
    i32.store
    local.get 0
    i32.const 2
    i32.store offset=44
    local.get 0
    i64.const 2
    i64.store offset=20 align=4
    local.get 0
    i32.const 1048876
    i32.store offset=16
    local.get 0
    local.get 0
    i32.const 56
    i32.add
    i32.store offset=48
    local.get 0
    local.get 0
    i32.const 8
    i32.add
    i32.store offset=40
    local.get 0
    local.get 0
    i32.const 40
    i32.add
    i32.store offset=32
    local.get 0
    i32.const 16
    i32.add
    i32.const 1048916
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN3std9panicking11begin_panic17h445db074890b5b34E (type 8) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    local.get 1
    i32.store offset=12
    local.get 3
    local.get 0
    i32.store offset=8
    local.get 3
    i32.const 8
    i32.add
    i32.const 1049884
    i32.const 0
    local.get 2
    call $_ZN3std9panicking20rust_panic_with_hook17h3f88564dfeb012ebE
    unreachable)
  (func $_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h5336016dc9200794E (type 1) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.load
    i32.store offset=12
    local.get 2
    i32.const 12
    i32.add
    local.get 1
    call $_ZN3std4sync4once4Once9call_once28_$u7b$$u7b$closure$u7d$$u7d$17he0ac5916c2b898e8E
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func $_ZN3std4sync4once4Once9call_once28_$u7b$$u7b$closure$u7d$$u7d$17he0ac5916c2b898e8E (type 1) (param i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    local.get 0
    i32.load
    local.tee 0
    i32.load8_u
    local.set 2
    local.get 0
    i32.const 0
    i32.store8
    block  ;; label = @1
      local.get 2
      i32.const 1
      i32.and
      i32.eqz
      br_if 0 (;@1;)
      i32.const 1
      local.set 3
      block  ;; label = @2
        block  ;; label = @3
          loop  ;; label = @4
            i32.const 0
            i32.load8_u offset=1056201
            br_if 1 (;@3;)
            i32.const 0
            i32.load offset=1055704
            local.set 4
            i32.const 0
            local.get 3
            i32.const 10
            i32.eq
            i32.store offset=1055704
            i32.const 0
            i32.const 0
            i32.store8 offset=1056201
            block  ;; label = @5
              local.get 4
              i32.eqz
              br_if 0 (;@5;)
              local.get 4
              i32.const 1
              i32.eq
              br_if 3 (;@2;)
              local.get 4
              i32.load
              local.tee 5
              local.get 4
              i32.load offset=8
              local.tee 2
              i32.const 3
              i32.shl
              i32.add
              local.set 6
              local.get 4
              i32.load offset=4
              local.set 7
              local.get 5
              local.set 0
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 5
                  local.set 0
                  block  ;; label = @8
                    loop  ;; label = @9
                      local.get 0
                      i32.load
                      local.tee 2
                      i32.eqz
                      br_if 1 (;@8;)
                      local.get 2
                      local.get 0
                      i32.const 4
                      i32.add
                      i32.load
                      call $_ZN83_$LT$alloc..boxed..Box$LT$F$GT$$u20$as$u20$core..ops..function..FnOnce$LT$A$GT$$GT$9call_once17hb8ffca9b0ba08c47E
                      local.get 0
                      i32.const 8
                      i32.add
                      local.tee 0
                      local.get 6
                      i32.ne
                      br_if 0 (;@9;)
                      br 3 (;@6;)
                    end
                  end
                  local.get 0
                  i32.const 8
                  i32.add
                  local.set 0
                end
                local.get 0
                local.get 6
                i32.eq
                br_if 0 (;@6;)
                loop  ;; label = @7
                  local.get 0
                  i32.load
                  local.tee 2
                  i32.eqz
                  br_if 1 (;@6;)
                  local.get 2
                  local.get 0
                  i32.const 4
                  i32.add
                  i32.load
                  local.tee 8
                  i32.load
                  call_indirect (type 0)
                  block  ;; label = @8
                    local.get 8
                    i32.load offset=4
                    local.tee 9
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 2
                    local.get 9
                    local.get 8
                    i32.load offset=8
                    call $__rust_dealloc
                  end
                  local.get 0
                  i32.const 8
                  i32.add
                  local.tee 0
                  local.get 6
                  i32.ne
                  br_if 0 (;@7;)
                end
              end
              block  ;; label = @6
                local.get 7
                i32.eqz
                br_if 0 (;@6;)
                local.get 5
                local.get 7
                i32.const 3
                i32.shl
                i32.const 4
                call $__rust_dealloc
              end
              local.get 4
              i32.const 12
              i32.const 4
              call $__rust_dealloc
            end
            local.get 3
            local.get 3
            i32.const 10
            i32.lt_u
            local.tee 0
            i32.add
            local.set 3
            local.get 0
            br_if 0 (;@4;)
          end
          return
        end
        i32.const 1050084
        i32.const 32
        i32.const 1050068
        call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
        unreachable
      end
      i32.const 1049700
      i32.const 31
      i32.const 1049684
      call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
      unreachable
    end
    i32.const 1048848
    call $_ZN4core9panicking5panic17h62fdcfa056e70982E
    unreachable)
  (func $_ZN4core3ptr18real_drop_in_place17h0372efb5da71b90dE (type 0) (param i32))
  (func $_ZN4core3ptr18real_drop_in_place17h41544376cef8a15bE (type 0) (param i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=4
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.load
      local.get 1
      i32.const 1
      call $__rust_dealloc
    end)
  (func $_ZN4core3ptr18real_drop_in_place17h7615a9d826f89d4cE (type 0) (param i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=4
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 8
      i32.add
      i32.load
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 0
      i32.const 1
      call $__rust_dealloc
    end)
  (func $_ZN4core3ptr18real_drop_in_place17hc5a29f899ad0c715E (type 0) (param i32))
  (func $_ZN4core6option15Option$LT$T$GT$6unwrap17hcab2ea581bad05cdE (type 2) (param i32) (result i32)
    block  ;; label = @1
      local.get 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      return
    end
    i32.const 1048848
    call $_ZN4core9panicking5panic17h62fdcfa056e70982E
    unreachable)
  (func $_ZN4core6option15Option$LT$T$GT$6unwrap17hfc66b6bb66dd1900E (type 2) (param i32) (result i32)
    block  ;; label = @1
      local.get 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      return
    end
    i32.const 1048848
    call $_ZN4core9panicking5panic17h62fdcfa056e70982E
    unreachable)
  (func $_ZN4core6result13unwrap_failed17h5ea5d9257702b76bE (type 1) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 43
    i32.store offset=12
    local.get 2
    i32.const 1048932
    i32.store offset=8
    local.get 2
    local.get 1
    i32.store8 offset=20
    local.get 2
    local.get 0
    i32.store offset=16
    local.get 2
    i32.const 60
    i32.add
    i32.const 3
    i32.store
    local.get 2
    i32.const 44
    i32.add
    i32.const 2
    i32.store
    local.get 2
    i32.const 2
    i32.store offset=52
    local.get 2
    i64.const 2
    i64.store offset=28 align=4
    local.get 2
    i32.const 1048876
    i32.store offset=24
    local.get 2
    local.get 2
    i32.const 16
    i32.add
    i32.store offset=56
    local.get 2
    local.get 2
    i32.const 8
    i32.add
    i32.store offset=48
    local.get 2
    local.get 2
    i32.const 48
    i32.add
    i32.store offset=40
    local.get 2
    i32.const 24
    i32.add
    i32.const 1048916
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN82_$LT$std..sys_common..poison..PoisonError$LT$T$GT$$u20$as$u20$core..fmt..Debug$GT$3fmt17h28d7bea72dfee2beE (type 4) (param i32 i32) (result i32)
    i32.const 1049731
    i32.const 25
    local.get 1
    call $_ZN40_$LT$str$u20$as$u20$core..fmt..Debug$GT$3fmt17h5b1f8bd45e1c8428E)
  (func $_ZN4core6result13unwrap_failed17h6b8b1582db93cfa3E (type 0) (param i32)
    (local i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i32.const 47
    i32.store offset=4
    local.get 1
    i32.const 1049271
    i32.store
    local.get 1
    i32.const 44
    i32.add
    i32.const 4
    i32.store
    local.get 1
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 1
    local.get 0
    i32.store offset=40
    local.get 1
    i32.const 2
    i32.store offset=36
    local.get 1
    i64.const 2
    i64.store offset=12 align=4
    local.get 1
    i32.const 1048876
    i32.store offset=8
    local.get 1
    local.get 1
    i32.store offset=32
    local.get 1
    local.get 1
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 1
    i32.const 8
    i32.add
    i32.const 1048916
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN62_$LT$std..ffi..c_str..NulError$u20$as$u20$core..fmt..Debug$GT$3fmt17h78d171950c7b08abE (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.const 1049936
    i32.const 8
    call $_ZN4core3fmt9Formatter11debug_tuple17h3d2072b8cea76c89E
    local.get 2
    local.get 0
    i32.store offset=12
    local.get 2
    local.get 2
    i32.const 12
    i32.add
    i32.const 1049060
    call $_ZN4core3fmt8builders10DebugTuple5field17hc049fbe3a16238abE
    drop
    local.get 2
    local.get 0
    i32.const 4
    i32.add
    i32.store offset=12
    local.get 2
    local.get 2
    i32.const 12
    i32.add
    i32.const 1049944
    call $_ZN4core3fmt8builders10DebugTuple5field17hc049fbe3a16238abE
    drop
    local.get 2
    call $_ZN4core3fmt8builders10DebugTuple6finish17h9e783d267307e9eaE
    local.set 0
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func $_ZN4core6result13unwrap_failed17h6d179f84b957db36E (type 5)
    (local i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 0
    global.set 0
    local.get 0
    i32.const 24
    i32.store offset=12
    local.get 0
    i32.const 1048632
    i32.store offset=8
    local.get 0
    i32.const 52
    i32.add
    i32.const 5
    i32.store
    local.get 0
    i32.const 36
    i32.add
    i32.const 2
    i32.store
    local.get 0
    i32.const 2
    i32.store offset=44
    local.get 0
    i64.const 2
    i64.store offset=20 align=4
    local.get 0
    i32.const 1048876
    i32.store offset=16
    local.get 0
    local.get 0
    i32.const 56
    i32.add
    i32.store offset=48
    local.get 0
    local.get 0
    i32.const 8
    i32.add
    i32.store offset=40
    local.get 0
    local.get 0
    i32.const 40
    i32.add
    i32.store offset=32
    local.get 0
    i32.const 16
    i32.add
    i32.const 1048916
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN3std9panicking15begin_panic_fmt17h4c661c5fa0dd4230E (type 1) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 32
    i32.add
    local.get 1
    i32.load
    local.get 1
    i32.load offset=4
    local.get 1
    i32.load offset=8
    local.get 1
    i32.load offset=12
    call $_ZN4core5panic8Location20internal_constructor17h55a9de981ef3aeb7E
    local.get 2
    i32.const 20
    i32.add
    local.get 2
    i32.const 40
    i32.add
    i64.load
    i64.store align=4
    local.get 2
    local.get 0
    i32.store offset=8
    local.get 2
    i32.const 1048740
    i32.store offset=4
    local.get 2
    i32.const 1048740
    i32.store
    local.get 2
    local.get 2
    i64.load offset=32
    i64.store offset=12 align=4
    local.get 2
    call $_ZN3std9panicking18continue_panic_fmt17he2d0fc2a935878afE
    unreachable)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17h6932f8aa37f23f2cE (type 4) (param i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    i32.load
    local.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 128
        i32.ge_u
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 0
          i32.load offset=8
          local.tee 3
          local.get 0
          i32.load offset=4
          i32.ne
          br_if 0 (;@3;)
          local.get 0
          i32.const 1
          call $_ZN5alloc3vec12Vec$LT$T$GT$7reserve17h03d7aeb8e635e142E
          local.get 0
          i32.const 8
          i32.add
          i32.load
          local.set 3
        end
        local.get 0
        i32.load
        local.get 3
        i32.add
        local.get 1
        i32.store8
        local.get 0
        i32.const 8
        i32.add
        local.tee 0
        local.get 0
        i32.load
        i32.const 1
        i32.add
        i32.store
        br 1 (;@1;)
      end
      local.get 2
      i32.const 0
      i32.store offset=12
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.const 2048
          i32.ge_u
          br_if 0 (;@3;)
          local.get 2
          local.get 1
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=13
          local.get 2
          local.get 1
          i32.const 6
          i32.shr_u
          i32.const 31
          i32.and
          i32.const 192
          i32.or
          i32.store8 offset=12
          i32.const 2
          local.set 1
          br 1 (;@2;)
        end
        block  ;; label = @3
          local.get 1
          i32.const 65535
          i32.gt_u
          br_if 0 (;@3;)
          local.get 2
          local.get 1
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=14
          local.get 2
          local.get 1
          i32.const 6
          i32.shr_u
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=13
          local.get 2
          local.get 1
          i32.const 12
          i32.shr_u
          i32.const 15
          i32.and
          i32.const 224
          i32.or
          i32.store8 offset=12
          i32.const 3
          local.set 1
          br 1 (;@2;)
        end
        local.get 2
        local.get 1
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=15
        local.get 2
        local.get 1
        i32.const 18
        i32.shr_u
        i32.const 240
        i32.or
        i32.store8 offset=12
        local.get 2
        local.get 1
        i32.const 6
        i32.shr_u
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=14
        local.get 2
        local.get 1
        i32.const 12
        i32.shr_u
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=13
        i32.const 4
        local.set 1
      end
      local.get 0
      local.get 1
      call $_ZN5alloc3vec12Vec$LT$T$GT$7reserve17h03d7aeb8e635e142E
      local.get 0
      local.get 0
      i32.load offset=8
      local.tee 3
      local.get 1
      i32.add
      i32.store offset=8
      local.get 3
      local.get 0
      i32.load
      i32.add
      local.get 2
      i32.const 12
      i32.add
      local.get 1
      call $memcpy
      drop
    end
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    i32.const 0)
  (func $_ZN5alloc3vec12Vec$LT$T$GT$7reserve17h03d7aeb8e635e142E (type 1) (param i32 i32)
    (local i32 i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=4
      local.tee 2
      local.get 0
      i32.load offset=8
      local.tee 3
      i32.sub
      local.get 1
      i32.ge_u
      br_if 0 (;@1;)
      block  ;; label = @2
        block  ;; label = @3
          local.get 3
          local.get 1
          i32.add
          local.tee 1
          local.get 3
          i32.lt_u
          br_if 0 (;@3;)
          local.get 2
          i32.const 1
          i32.shl
          local.tee 3
          local.get 1
          local.get 1
          local.get 3
          i32.lt_u
          select
          local.tee 1
          i32.const 0
          i32.lt_s
          br_if 0 (;@3;)
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.eqz
              br_if 0 (;@5;)
              local.get 0
              i32.load
              local.get 2
              i32.const 1
              local.get 1
              call $__rust_realloc
              local.tee 2
              i32.eqz
              br_if 1 (;@4;)
              br 3 (;@2;)
            end
            local.get 1
            i32.const 1
            call $__rust_alloc
            local.tee 2
            br_if 2 (;@2;)
          end
          local.get 1
          i32.const 1
          call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
          unreachable
        end
        call $_ZN5alloc7raw_vec17capacity_overflow17hd685e916963b651dE
        unreachable
      end
      local.get 0
      local.get 2
      i32.store
      local.get 0
      i32.const 4
      i32.add
      local.get 1
      i32.store
    end)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17hfbe1b10105aef519E (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.load
    i32.store offset=4
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 1
    i64.load align=4
    i64.store offset=8
    local.get 2
    i32.const 4
    i32.add
    i32.const 1048576
    local.get 2
    i32.const 8
    i32.add
    call $_ZN4core3fmt5write17h8cfd01c67a4a46c9E
    local.set 1
    local.get 2
    i32.const 32
    i32.add
    global.set 0
    local.get 1)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17h2ea757743c87a7b0E (type 3) (param i32 i32 i32) (result i32)
    (local i32)
    local.get 0
    i32.load
    local.tee 0
    local.get 2
    call $_ZN5alloc3vec12Vec$LT$T$GT$7reserve17h03d7aeb8e635e142E
    local.get 0
    local.get 0
    i32.load offset=8
    local.tee 3
    local.get 2
    i32.add
    i32.store offset=8
    local.get 3
    local.get 0
    i32.load
    i32.add
    local.get 1
    local.get 2
    call $memcpy
    drop
    i32.const 0)
  (func $_ZN5alloc4sync12Arc$LT$T$GT$9drop_slow17h6d4c8a2e1ae6d42bE (type 0) (param i32)
    (local i32 i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 1
      i32.const 16
      i32.add
      i32.load
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      i32.const 0
      i32.store8
      local.get 1
      i32.const 20
      i32.add
      i32.load
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 16
      i32.add
      i32.load
      local.get 2
      i32.const 1
      call $__rust_dealloc
    end
    local.get 1
    i32.const 28
    i32.add
    i32.load
    i32.const 1
    i32.const 1
    call $__rust_dealloc
    local.get 0
    i32.load
    local.tee 1
    local.get 1
    i32.load offset=4
    local.tee 1
    i32.const -1
    i32.add
    i32.store offset=4
    block  ;; label = @1
      local.get 1
      i32.const 1
      i32.ne
      br_if 0 (;@1;)
      local.get 0
      i32.load
      i32.const 48
      i32.const 8
      call $__rust_dealloc
    end)
  (func $_ZN83_$LT$alloc..boxed..Box$LT$F$GT$$u20$as$u20$core..ops..function..FnOnce$LT$A$GT$$GT$9call_once17hb8ffca9b0ba08c47E (type 1) (param i32 i32)
    (local i32 i32 i32)
    global.get 0
    local.tee 2
    local.set 3
    local.get 2
    local.get 1
    i32.load offset=4
    local.tee 4
    i32.const 15
    i32.add
    i32.const -16
    i32.and
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    local.get 4
    call $memcpy
    local.get 1
    i32.load offset=12
    call_indirect (type 0)
    block  ;; label = @1
      local.get 4
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 4
      local.get 1
      i32.load offset=8
      call $__rust_dealloc
    end
    local.get 3
    global.set 0)
  (func $_ZN3std10sys_common11thread_info10ThreadInfo4with28_$u7b$$u7b$closure$u7d$$u7d$17h7ac7c659d1fa9b1aE (type 2) (param i32) (result i32)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 0
            i32.load
            local.tee 2
            i32.const 0
            i32.lt_s
            br_if 0 (;@4;)
            local.get 2
            i32.const 2147483647
            i32.eq
            br_if 0 (;@4;)
            local.get 0
            local.get 2
            i32.store
            block  ;; label = @5
              local.get 0
              i32.load offset=4
              local.tee 3
              br_if 0 (;@5;)
              local.get 1
              i32.const 0
              i32.store
              local.get 1
              call $_ZN3std6thread6Thread3new17h0a174627a8dd7c63E
              local.set 3
              local.get 0
              i32.load
              br_if 2 (;@3;)
              local.get 0
              i32.const -1
              i32.store
              block  ;; label = @6
                local.get 0
                i32.const 4
                i32.add
                local.tee 4
                i32.load
                local.tee 2
                i32.eqz
                br_if 0 (;@6;)
                local.get 2
                local.get 2
                i32.load
                local.tee 5
                i32.const -1
                i32.add
                i32.store
                local.get 5
                i32.const 1
                i32.ne
                br_if 0 (;@6;)
                local.get 0
                i32.const 4
                i32.add
                call $_ZN5alloc4sync12Arc$LT$T$GT$9drop_slow17h6d4c8a2e1ae6d42bE
              end
              local.get 4
              local.get 3
              i32.store
              local.get 0
              local.get 0
              i32.load
              i32.const 1
              i32.add
              local.tee 2
              i32.store
            end
            local.get 2
            br_if 1 (;@3;)
            local.get 0
            i32.const -1
            i32.store
            local.get 3
            i32.eqz
            br_if 2 (;@2;)
            local.get 3
            local.get 3
            i32.load
            local.tee 2
            i32.const 1
            i32.add
            i32.store
            local.get 2
            i32.const -1
            i32.le_s
            br_if 3 (;@1;)
            local.get 0
            local.get 0
            i32.load
            i32.const 1
            i32.add
            i32.store
            local.get 1
            i32.const 16
            i32.add
            global.set 0
            local.get 3
            return
          end
          call $_ZN4core6result13unwrap_failed17h6d179f84b957db36E
          unreachable
        end
        call $_ZN4core6result13unwrap_failed17h1984afaa28d103e8E
        unreachable
      end
      i32.const 1048848
      call $_ZN4core9panicking5panic17h62fdcfa056e70982E
      unreachable
    end
    unreachable
    unreachable)
  (func $_ZN3std6thread4park17h5a4b125952deab3eE (type 5)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 96
    i32.sub
    local.tee 0
    global.set 0
    block  ;; label = @1
      i32.const 0
      i32.load offset=1055728
      i32.const 1
      i32.eq
      br_if 0 (;@1;)
      i32.const 0
      i64.const 1
      i64.store offset=1055728 align=4
      i32.const 0
      i32.const 0
      i32.store offset=1055736
    end
    i32.const 1055732
    call $_ZN3std10sys_common11thread_info10ThreadInfo4with28_$u7b$$u7b$closure$u7d$$u7d$17h7ac7c659d1fa9b1aE
    local.tee 1
    i32.const 0
    local.get 1
    i32.load offset=24
    local.tee 2
    local.get 2
    i32.const 2
    i32.eq
    select
    i32.store offset=24
    local.get 0
    local.get 1
    i32.store offset=8
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 2
                i32.const 2
                i32.eq
                br_if 0 (;@6;)
                local.get 0
                i32.load offset=8
                local.tee 1
                i32.const 28
                i32.add
                i32.load
                local.tee 2
                i32.load8_u
                br_if 1 (;@5;)
                local.get 2
                i32.const 1
                i32.store8
                i32.const 0
                local.set 3
                block  ;; label = @7
                  block  ;; label = @8
                    i32.const 0
                    i32.load offset=1056192
                    i32.const 1
                    i32.ne
                    br_if 0 (;@8;)
                    i32.const 0
                    i32.load offset=1056196
                    local.set 3
                    br 1 (;@7;)
                  end
                  i32.const 0
                  i64.const 1
                  i64.store offset=1056192
                end
                i32.const 0
                local.get 3
                i32.store offset=1056196
                local.get 1
                i32.const 32
                i32.add
                i32.load8_u
                br_if 2 (;@4;)
                local.get 1
                i32.const 24
                i32.add
                local.tee 2
                local.get 2
                i32.load
                local.tee 2
                i32.const 1
                local.get 2
                select
                i32.store
                local.get 2
                i32.eqz
                br_if 3 (;@3;)
                local.get 2
                i32.const 2
                i32.ne
                br_if 4 (;@2;)
                local.get 0
                i32.load offset=8
                i32.const 24
                i32.add
                local.tee 4
                i32.load
                local.set 2
                local.get 4
                i32.const 0
                i32.store
                local.get 0
                local.get 2
                i32.store offset=12
                local.get 2
                i32.const 2
                i32.ne
                br_if 5 (;@1;)
                block  ;; label = @7
                  local.get 3
                  br_if 0 (;@7;)
                  block  ;; label = @8
                    i32.const 0
                    i32.load offset=1056192
                    i32.const 1
                    i32.ne
                    br_if 0 (;@8;)
                    i32.const 0
                    i32.load offset=1056196
                    i32.eqz
                    br_if 1 (;@7;)
                    local.get 1
                    i32.const 32
                    i32.add
                    i32.const 1
                    i32.store8
                    br 1 (;@7;)
                  end
                  i32.const 0
                  i64.const 1
                  i64.store offset=1056192
                end
                local.get 1
                i32.const 28
                i32.add
                i32.load
                i32.const 0
                i32.store8
              end
              local.get 0
              i32.load offset=8
              local.tee 1
              local.get 1
              i32.load
              local.tee 1
              i32.const -1
              i32.add
              i32.store
              block  ;; label = @6
                local.get 1
                i32.const 1
                i32.ne
                br_if 0 (;@6;)
                local.get 0
                i32.const 8
                i32.add
                call $_ZN5alloc4sync12Arc$LT$T$GT$9drop_slow17h6d4c8a2e1ae6d42bE
              end
              local.get 0
              i32.const 96
              i32.add
              global.set 0
              return
            end
            i32.const 1050084
            i32.const 32
            i32.const 1050068
            call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
            unreachable
          end
          local.get 1
          i32.const 28
          i32.add
          local.get 3
          i32.const 0
          i32.ne
          call $_ZN4core6result13unwrap_failed17h5ea5d9257702b76bE
          unreachable
        end
        local.get 0
        i32.load offset=8
        i32.const 36
        i32.add
        local.tee 0
        local.get 1
        i32.const 28
        i32.add
        i32.load
        local.tee 1
        call $_ZN3std4sync7condvar7Condvar6verify17h1f39047be07bdb0cE
        local.get 0
        i32.load
        local.get 1
        call $_ZN3std10sys_common7condvar7Condvar4wait17h53625253e5401802E
        unreachable
      end
      i32.const 1049116
      i32.const 23
      i32.const 1049100
      call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
      unreachable
    end
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    i32.store offset=64
    local.get 0
    i32.const 1049140
    i32.store offset=68
    local.get 0
    i32.const 72
    i32.add
    i32.const 20
    i32.add
    i32.const 0
    i32.store
    local.get 0
    i32.const 40
    i32.add
    i32.const 20
    i32.add
    i32.const 6
    i32.store
    local.get 0
    i32.const 52
    i32.add
    i32.const 7
    i32.store
    local.get 0
    i32.const 16
    i32.add
    i32.const 20
    i32.add
    i32.const 3
    i32.store
    local.get 0
    i32.const 1048740
    i32.store offset=88
    local.get 0
    i64.const 1
    i64.store offset=76 align=4
    local.get 0
    i32.const 1049176
    i32.store offset=72
    local.get 0
    i32.const 7
    i32.store offset=44
    local.get 0
    i64.const 3
    i64.store offset=20 align=4
    local.get 0
    i32.const 1048760
    i32.store offset=16
    local.get 0
    local.get 0
    i32.const 72
    i32.add
    i32.store offset=56
    local.get 0
    local.get 0
    i32.const 68
    i32.add
    i32.store offset=48
    local.get 0
    local.get 0
    i32.const 64
    i32.add
    i32.store offset=40
    local.get 0
    local.get 0
    i32.const 40
    i32.add
    i32.store offset=32
    local.get 0
    i32.const 16
    i32.add
    i32.const 1049184
    call $_ZN3std9panicking15begin_panic_fmt17h4c661c5fa0dd4230E
    unreachable)
  (func $_ZN3std4sync7condvar7Condvar6verify17h1f39047be07bdb0cE (type 1) (param i32 i32)
    (local i32)
    local.get 0
    local.get 0
    i32.load offset=4
    local.tee 2
    local.get 1
    local.get 2
    select
    i32.store offset=4
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        local.get 1
        i32.ne
        br_if 1 (;@1;)
      end
      return
    end
    i32.const 1049408
    i32.const 54
    i32.const 1049392
    call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
    unreachable)
  (func $_ZN3std10sys_common7condvar7Condvar4wait17h53625253e5401802E (type 1) (param i32 i32)
    (local i32)
    local.get 2
    local.get 2
    call $_ZN3std3sys4wasm7condvar7Condvar4wait17h8b15799de3f3e652E
    unreachable)
  (func $_ZN3std6thread6Thread3new17h0a174627a8dd7c63E (type 2) (param i32) (result i32)
    (local i32 i32 i32 i32 i64)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 0
            i32.load
            local.tee 2
            i32.eqz
            br_if 0 (;@4;)
            local.get 1
            local.get 0
            i64.load offset=4 align=4
            i64.store offset=36 align=4
            local.get 1
            local.get 2
            i32.store offset=32
            local.get 1
            i32.const 16
            i32.add
            local.get 1
            i32.const 32
            i32.add
            call $_ZN5alloc6string104_$LT$impl$u20$core..convert..From$LT$alloc..string..String$GT$$u20$for$u20$alloc..vec..Vec$LT$u8$GT$$GT$4from17he2b8ded24866ca30E
            local.get 1
            i32.const 8
            i32.add
            i32.const 0
            local.get 1
            i32.load offset=16
            local.tee 0
            local.get 1
            i32.load offset=24
            call $_ZN4core5slice6memchr6memchr17h6bde110ce4c09380E
            local.get 1
            i32.load offset=8
            br_if 2 (;@2;)
            local.get 1
            i32.const 32
            i32.add
            i32.const 8
            i32.add
            local.get 1
            i32.const 16
            i32.add
            i32.const 8
            i32.add
            i32.load
            i32.store
            local.get 1
            local.get 1
            i64.load offset=16
            i64.store offset=32
            local.get 1
            local.get 1
            i32.const 32
            i32.add
            call $_ZN3std3ffi5c_str7CString18from_vec_unchecked17ha1322458045ed041E
            local.get 1
            i32.load offset=4
            local.set 3
            local.get 1
            i32.load
            local.set 4
            i32.const 0
            i32.load8_u offset=1056200
            br_if 1 (;@3;)
            br 3 (;@1;)
          end
          i32.const 0
          local.set 4
          i32.const 0
          i32.load8_u offset=1056200
          i32.eqz
          br_if 2 (;@1;)
        end
        i32.const 1050084
        i32.const 32
        i32.const 1050068
        call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
        unreachable
      end
      local.get 1
      i32.load offset=12
      local.set 2
      local.get 1
      i32.const 40
      i32.add
      local.get 1
      i64.load offset=20 align=4
      i64.store
      local.get 1
      local.get 0
      i32.store offset=36
      local.get 1
      local.get 2
      i32.store offset=32
      local.get 1
      i32.const 32
      i32.add
      call $_ZN4core6result13unwrap_failed17h6b8b1582db93cfa3E
      unreachable
    end
    i32.const 0
    i32.const 1
    i32.store8 offset=1056200
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            i32.const 0
            i64.load offset=1055696
            local.tee 5
            i64.const -1
            i64.eq
            br_if 0 (;@4;)
            i32.const 0
            local.get 5
            i64.const 1
            i64.add
            i64.store offset=1055696
            local.get 5
            i64.const 0
            i64.eq
            br_if 1 (;@3;)
            i32.const 0
            i32.const 0
            i32.store8 offset=1056200
            i32.const 1
            i32.const 1
            call $__rust_alloc
            local.tee 2
            i32.eqz
            br_if 2 (;@2;)
            local.get 2
            i32.const 0
            i32.store8
            i32.const 48
            i32.const 8
            call $__rust_alloc
            local.tee 0
            i32.eqz
            br_if 3 (;@1;)
            local.get 0
            i64.const 1
            i64.store offset=36 align=4
            local.get 0
            i32.const 0
            i32.store offset=24
            local.get 0
            local.get 3
            i32.store offset=20
            local.get 0
            local.get 4
            i32.store offset=16
            local.get 0
            local.get 5
            i64.store offset=8
            local.get 0
            i64.const 4294967297
            i64.store
            local.get 0
            local.get 2
            i64.extend_i32_u
            i64.store offset=28 align=4
            local.get 1
            i32.const 48
            i32.add
            global.set 0
            local.get 0
            return
          end
          i32.const 1049216
          i32.const 55
          i32.const 1049200
          call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
          unreachable
        end
        i32.const 1048848
        call $_ZN4core9panicking5panic17h62fdcfa056e70982E
        unreachable
      end
      i32.const 1
      i32.const 1
      call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
      unreachable
    end
    i32.const 48
    i32.const 8
    call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
    unreachable)
  (func $_ZN3std3ffi5c_str7CString18from_vec_unchecked17ha1322458045ed041E (type 1) (param i32 i32)
    (local i32 i32 i32 i32)
    block  ;; label = @1
      local.get 1
      i32.load offset=4
      local.tee 2
      local.get 1
      i32.load offset=8
      local.tee 3
      i32.ne
      br_if 0 (;@1;)
      block  ;; label = @2
        block  ;; label = @3
          local.get 3
          i32.const 1
          i32.add
          local.tee 2
          i32.const 0
          i32.lt_s
          br_if 0 (;@3;)
          local.get 2
          local.get 3
          i32.lt_u
          br_if 0 (;@3;)
          block  ;; label = @4
            block  ;; label = @5
              local.get 3
              i32.eqz
              br_if 0 (;@5;)
              local.get 1
              i32.load
              local.get 3
              i32.const 1
              local.get 2
              call $__rust_realloc
              local.tee 4
              i32.eqz
              br_if 1 (;@4;)
              br 3 (;@2;)
            end
            local.get 2
            i32.const 1
            call $__rust_alloc
            local.tee 4
            br_if 2 (;@2;)
          end
          local.get 2
          i32.const 1
          call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
          unreachable
        end
        call $_ZN5alloc7raw_vec17capacity_overflow17hd685e916963b651dE
        unreachable
      end
      local.get 1
      local.get 4
      i32.store
      local.get 1
      i32.const 4
      i32.add
      local.get 2
      i32.store
    end
    block  ;; label = @1
      local.get 3
      local.get 2
      i32.ne
      br_if 0 (;@1;)
      local.get 1
      i32.const 1
      call $_ZN5alloc3vec12Vec$LT$T$GT$7reserve17h03d7aeb8e635e142E
      local.get 1
      i32.const 4
      i32.add
      i32.load
      local.set 2
      local.get 1
      i32.const 8
      i32.add
      i32.load
      local.set 3
    end
    local.get 1
    i32.const 8
    i32.add
    local.get 3
    i32.const 1
    i32.add
    local.tee 4
    i32.store
    local.get 1
    i32.load
    local.tee 5
    local.get 3
    i32.add
    i32.const 0
    i32.store8
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          local.get 4
          i32.ne
          br_if 0 (;@3;)
          local.get 5
          local.set 1
          local.get 2
          local.set 4
          br 1 (;@2;)
        end
        local.get 2
        local.get 4
        i32.lt_u
        br_if 1 (;@1;)
        block  ;; label = @3
          local.get 4
          i32.eqz
          br_if 0 (;@3;)
          local.get 5
          local.get 2
          i32.const 1
          local.get 4
          call $__rust_realloc
          local.tee 1
          br_if 1 (;@2;)
          local.get 4
          i32.const 1
          call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
          unreachable
        end
        i32.const 0
        local.set 4
        i32.const 1
        local.set 1
        local.get 2
        i32.eqz
        br_if 0 (;@2;)
        local.get 5
        local.get 2
        i32.const 1
        call $__rust_dealloc
      end
      local.get 0
      local.get 4
      i32.store offset=4
      local.get 0
      local.get 1
      i32.store
      return
    end
    i32.const 1049036
    call $_ZN4core9panicking5panic17h62fdcfa056e70982E
    unreachable)
  (func $_ZN3std4sync4once4Once10call_inner17h92989b0c916494d1E (type 9) (param i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 4
    global.set 0
    local.get 4
    i32.const 2
    i32.or
    local.set 5
    local.get 0
    i32.load
    local.set 6
    local.get 4
    i32.const 8
    i32.add
    local.set 7
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            loop  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 6
                  local.tee 8
                  i32.eqz
                  br_if 0 (;@7;)
                  block  ;; label = @8
                    local.get 8
                    i32.const 1
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 8
                    i32.const 3
                    i32.eq
                    br_if 5 (;@3;)
                    local.get 8
                    i32.const 3
                    i32.and
                    i32.const 2
                    i32.ne
                    br_if 6 (;@2;)
                    block  ;; label = @9
                      i32.const 0
                      i32.load offset=1055728
                      i32.const 1
                      i32.eq
                      br_if 0 (;@9;)
                      i32.const 0
                      i64.const 1
                      i64.store offset=1055728 align=4
                      i32.const 0
                      i32.const 0
                      i32.store offset=1055736
                    end
                    i32.const 1055732
                    call $_ZN3std10sys_common11thread_info10ThreadInfo4with28_$u7b$$u7b$closure$u7d$$u7d$17h7ac7c659d1fa9b1aE
                    local.set 6
                    local.get 7
                    i32.const 0
                    i32.store8
                    local.get 4
                    local.get 6
                    i32.store
                    local.get 4
                    i32.const 0
                    i32.store offset=4
                    local.get 8
                    local.set 6
                    loop  ;; label = @9
                      local.get 6
                      i32.const 3
                      i32.and
                      i32.const 2
                      i32.ne
                      br_if 3 (;@6;)
                      local.get 0
                      local.get 5
                      local.get 0
                      i32.load
                      local.tee 8
                      local.get 8
                      local.get 6
                      i32.eq
                      select
                      i32.store
                      local.get 4
                      local.get 6
                      i32.const -4
                      i32.and
                      i32.store offset=4
                      local.get 8
                      local.get 6
                      i32.ne
                      local.set 9
                      local.get 8
                      local.set 6
                      local.get 9
                      br_if 0 (;@9;)
                    end
                    block  ;; label = @9
                      local.get 7
                      i32.load8_u
                      br_if 0 (;@9;)
                      loop  ;; label = @10
                        call $_ZN3std6thread4park17h5a4b125952deab3eE
                        local.get 7
                        i32.load8_u
                        i32.eqz
                        br_if 0 (;@10;)
                      end
                    end
                    local.get 0
                    i32.load
                    local.set 6
                    local.get 4
                    i32.load
                    local.tee 8
                    i32.eqz
                    br_if 3 (;@5;)
                    local.get 8
                    local.get 8
                    i32.load
                    local.tee 9
                    i32.const -1
                    i32.add
                    i32.store
                    local.get 9
                    i32.const 1
                    i32.ne
                    br_if 3 (;@5;)
                    local.get 4
                    call $_ZN5alloc4sync12Arc$LT$T$GT$9drop_slow17h6d4c8a2e1ae6d42bE
                    br 3 (;@5;)
                  end
                  local.get 1
                  i32.eqz
                  br_if 6 (;@1;)
                end
                local.get 0
                i32.const 2
                local.get 0
                i32.load
                local.tee 6
                local.get 6
                local.get 8
                i32.eq
                select
                i32.store
                local.get 6
                local.get 8
                i32.ne
                br_if 1 (;@5;)
                br 2 (;@4;)
              end
              local.get 4
              i32.load
              local.tee 8
              i32.eqz
              br_if 0 (;@5;)
              local.get 8
              local.get 8
              i32.load
              local.tee 9
              i32.const -1
              i32.add
              i32.store
              local.get 9
              i32.const 1
              i32.ne
              br_if 0 (;@5;)
              local.get 4
              call $_ZN5alloc4sync12Arc$LT$T$GT$9drop_slow17h6d4c8a2e1ae6d42bE
              br 0 (;@5;)
            end
          end
          local.get 4
          local.get 0
          i32.store
          local.get 2
          local.get 8
          i32.const 1
          i32.eq
          local.get 3
          i32.load offset=12
          call_indirect (type 1)
          local.get 4
          i32.const 0
          i32.store8 offset=4
          local.get 4
          call $_ZN65_$LT$std..sync..once..Finish$u20$as$u20$core..ops..drop..Drop$GT$4drop17h33951fe4fe3deca5E
        end
        local.get 4
        i32.const 16
        i32.add
        global.set 0
        return
      end
      i32.const 1049524
      i32.const 47
      i32.const 1049508
      call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
      unreachable
    end
    i32.const 1049588
    i32.const 42
    i32.const 1049572
    call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
    unreachable)
  (func $_ZN65_$LT$std..sync..once..Finish$u20$as$u20$core..ops..drop..Drop$GT$4drop17h33951fe4fe3deca5E (type 0) (param i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    i32.load
    local.tee 2
    i32.load
    local.set 3
    local.get 2
    i32.const 1
    i32.const 3
    local.get 0
    i32.load8_u offset=4
    select
    i32.store
    local.get 1
    local.get 3
    i32.const 3
    i32.and
    local.tee 0
    i32.store offset=12
    block  ;; label = @1
      local.get 0
      i32.const 2
      i32.ne
      br_if 0 (;@1;)
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 3
                i32.const -4
                i32.and
                local.tee 0
                i32.eqz
                br_if 0 (;@6;)
                loop  ;; label = @7
                  local.get 0
                  i32.load offset=4
                  local.set 2
                  local.get 0
                  i32.load
                  local.set 3
                  local.get 0
                  i32.const 0
                  i32.store
                  local.get 3
                  i32.eqz
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const 1
                  i32.store8 offset=8
                  local.get 1
                  local.get 3
                  i32.store offset=16
                  local.get 3
                  i32.const 24
                  i32.add
                  local.tee 3
                  i32.load
                  local.set 0
                  local.get 3
                  i32.const 2
                  i32.store
                  block  ;; label = @8
                    local.get 0
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 0
                    i32.const 2
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 0
                    i32.const 1
                    i32.ne
                    br_if 4 (;@4;)
                    local.get 1
                    i32.load offset=16
                    local.tee 0
                    i32.const 28
                    i32.add
                    local.tee 3
                    i32.load
                    local.tee 4
                    i32.load8_u
                    br_if 5 (;@3;)
                    local.get 4
                    i32.const 1
                    i32.store8
                    block  ;; label = @9
                      block  ;; label = @10
                        i32.const 0
                        i32.load offset=1056192
                        i32.const 1
                        i32.ne
                        br_if 0 (;@10;)
                        i32.const 0
                        i32.load offset=1056196
                        local.set 4
                        br 1 (;@9;)
                      end
                      i32.const 0
                      local.set 4
                      i32.const 0
                      i64.const 1
                      i64.store offset=1056192
                    end
                    i32.const 0
                    local.get 4
                    i32.store offset=1056196
                    local.get 0
                    i32.const 32
                    i32.add
                    i32.load8_u
                    br_if 6 (;@2;)
                    local.get 3
                    i32.load
                    i32.const 0
                    i32.store8
                  end
                  local.get 1
                  i32.load offset=16
                  local.tee 0
                  local.get 0
                  i32.load
                  local.tee 0
                  i32.const -1
                  i32.add
                  i32.store
                  block  ;; label = @8
                    local.get 0
                    i32.const 1
                    i32.ne
                    br_if 0 (;@8;)
                    local.get 1
                    i32.const 16
                    i32.add
                    call $_ZN5alloc4sync12Arc$LT$T$GT$9drop_slow17h6d4c8a2e1ae6d42bE
                  end
                  local.get 2
                  local.set 0
                  local.get 2
                  br_if 0 (;@7;)
                end
              end
              local.get 1
              i32.const 64
              i32.add
              global.set 0
              return
            end
            i32.const 1048848
            call $_ZN4core9panicking5panic17h62fdcfa056e70982E
            unreachable
          end
          i32.const 1049336
          i32.const 28
          i32.const 1049320
          call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
          unreachable
        end
        i32.const 1050084
        i32.const 32
        i32.const 1050068
        call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
        unreachable
      end
      local.get 0
      i32.const 28
      i32.add
      local.get 4
      i32.const 0
      i32.ne
      call $_ZN4core6result13unwrap_failed17h5ea5d9257702b76bE
      unreachable
    end
    local.get 1
    local.get 1
    i32.const 12
    i32.add
    i32.store offset=56
    local.get 1
    i32.const 52
    i32.add
    i32.const 7
    i32.store
    local.get 1
    i32.const 36
    i32.add
    i32.const 2
    i32.store
    local.get 1
    i32.const 7
    i32.store offset=44
    local.get 1
    i32.const 1049140
    i32.store offset=60
    local.get 1
    i64.const 3
    i64.store offset=20 align=4
    local.get 1
    i32.const 1048716
    i32.store offset=16
    local.get 1
    local.get 1
    i32.const 60
    i32.add
    i32.store offset=48
    local.get 1
    local.get 1
    i32.const 56
    i32.add
    i32.store offset=40
    local.get 1
    local.get 1
    i32.const 40
    i32.add
    i32.store offset=32
    local.get 1
    i32.const 16
    i32.add
    i32.const 1049632
    call $_ZN3std9panicking15begin_panic_fmt17h4c661c5fa0dd4230E
    unreachable)
  (func $_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17heab2a26d47adae2dE (type 4) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.load offset=12
    call_indirect (type 2))
  (func $_ZN3std3sys4wasm7condvar7Condvar4wait17h8b15799de3f3e652E (type 1) (param i32 i32)
    i32.const 1050008
    i32.const 29
    i32.const 1049992
    call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
    unreachable)
  (func $_ZN76_$LT$std..sys_common..thread_local..Key$u20$as$u20$core..ops..drop..Drop$GT$4drop17he3d36d5f1169eee9E (type 0) (param i32))
  (func $_ZN3std5alloc24default_alloc_error_hook17h4dee2a5fdd1dba75E (type 1) (param i32 i32))
  (func $rust_oom (type 1) (param i32 i32)
    (local i32)
    local.get 0
    local.get 1
    i32.const 0
    i32.load offset=1055712
    local.tee 2
    i32.const 8
    local.get 2
    select
    call_indirect (type 1)
    unreachable
    unreachable)
  (func $__rdl_alloc (type 4) (param i32 i32) (result i32)
    block  ;; label = @1
      i32.const 1055740
      call $_ZN8dlmalloc8dlmalloc8Dlmalloc16malloc_alignment17h1f8103281a1655e6E
      local.get 1
      i32.ge_u
      br_if 0 (;@1;)
      i32.const 1055740
      local.get 1
      local.get 0
      call $_ZN8dlmalloc8dlmalloc8Dlmalloc8memalign17hfb05805b97584ef4E
      return
    end
    i32.const 1055740
    local.get 0
    call $_ZN8dlmalloc8dlmalloc8Dlmalloc6malloc17h7ef7d8a98d4afe8dE)
  (func $__rdl_dealloc (type 8) (param i32 i32 i32)
    i32.const 1055740
    local.get 0
    call $_ZN8dlmalloc8dlmalloc8Dlmalloc4free17he25d64d2cc54b15fE)
  (func $__rdl_realloc (type 10) (param i32 i32 i32 i32) (result i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            i32.const 1055740
            call $_ZN8dlmalloc8dlmalloc8Dlmalloc16malloc_alignment17h1f8103281a1655e6E
            local.get 2
            i32.ge_u
            br_if 0 (;@4;)
            i32.const 1055740
            call $_ZN8dlmalloc8dlmalloc8Dlmalloc16malloc_alignment17h1f8103281a1655e6E
            local.get 2
            i32.ge_u
            br_if 1 (;@3;)
            i32.const 1055740
            local.get 2
            local.get 3
            call $_ZN8dlmalloc8dlmalloc8Dlmalloc8memalign17hfb05805b97584ef4E
            local.tee 2
            i32.eqz
            br_if 2 (;@2;)
            br 3 (;@1;)
          end
          i32.const 1055740
          local.get 0
          local.get 3
          call $_ZN8dlmalloc8dlmalloc8Dlmalloc7realloc17h60cd22bbd1e08b5cE
          return
        end
        i32.const 1055740
        local.get 3
        call $_ZN8dlmalloc8dlmalloc8Dlmalloc6malloc17h7ef7d8a98d4afe8dE
        local.tee 2
        br_if 1 (;@1;)
      end
      i32.const 0
      return
    end
    local.get 2
    local.get 0
    local.get 3
    local.get 1
    local.get 1
    local.get 3
    i32.gt_u
    select
    call $memcpy
    local.set 2
    i32.const 1055740
    local.get 0
    call $_ZN8dlmalloc8dlmalloc8Dlmalloc4free17he25d64d2cc54b15fE
    local.get 2)
  (func $_ZN3std9panicking3try7do_call17hb9a857ad47fc7875E (type 0) (param i32)
    (local i32)
    local.get 0
    local.get 0
    i32.load
    local.tee 1
    i32.load
    local.get 1
    i32.load offset=4
    call $_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17heab2a26d47adae2dE
    i32.store)
  (func $rust_begin_unwind (type 0) (param i32)
    local.get 0
    call $_ZN3std9panicking18continue_panic_fmt17he2d0fc2a935878afE
    unreachable)
  (func $_ZN3std9panicking18continue_panic_fmt17he2d0fc2a935878afE (type 0) (param i32)
    (local i32 i32 i32 i64 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    call $_ZN4core5panic9PanicInfo8location17h959e627e2fd51988E
    call $_ZN4core6option15Option$LT$T$GT$6unwrap17hfc66b6bb66dd1900E
    local.set 2
    local.get 0
    call $_ZN4core5panic9PanicInfo7message17hbe01febe97d5f595E
    call $_ZN4core6option15Option$LT$T$GT$6unwrap17hcab2ea581bad05cdE
    local.set 3
    local.get 1
    i32.const 8
    i32.add
    local.get 2
    call $_ZN4core5panic8Location4file17h9e415b81e07881c2E
    local.get 1
    i64.load offset=8
    local.set 4
    local.get 2
    call $_ZN4core5panic8Location4line17h4caf5525f6ea07a3E
    local.set 5
    local.get 1
    local.get 2
    call $_ZN4core5panic8Location6column17h3de733c09fa199ffE
    i32.store offset=28
    local.get 1
    local.get 5
    i32.store offset=24
    local.get 1
    local.get 4
    i64.store offset=16
    local.get 1
    i32.const 0
    i32.store offset=36
    local.get 1
    local.get 3
    i32.store offset=32
    local.get 1
    i32.const 32
    i32.add
    i32.const 1049848
    local.get 0
    call $_ZN4core5panic9PanicInfo7message17hbe01febe97d5f595E
    local.get 1
    i32.const 16
    i32.add
    call $_ZN3std9panicking20rust_panic_with_hook17h3f88564dfeb012ebE
    unreachable)
  (func $_ZN3std9panicking20rust_panic_with_hook17h3f88564dfeb012ebE (type 9) (param i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 4
    global.set 0
    i32.const 1
    local.set 5
    local.get 3
    i32.load offset=12
    local.set 6
    local.get 3
    i32.load offset=8
    local.set 7
    local.get 3
    i32.load offset=4
    local.set 8
    local.get 3
    i32.load
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            i32.const 0
            i32.load offset=1056192
            i32.const 1
            i32.ne
            br_if 0 (;@4;)
            i32.const 0
            i32.const 0
            i32.load offset=1056196
            i32.const 1
            i32.add
            local.tee 5
            i32.store offset=1056196
            local.get 5
            i32.const 3
            i32.lt_u
            br_if 1 (;@3;)
            br 2 (;@2;)
          end
          i32.const 0
          i64.const 4294967297
          i64.store offset=1056192
        end
        local.get 4
        i32.const 48
        i32.add
        local.get 3
        local.get 8
        local.get 7
        local.get 6
        call $_ZN4core5panic8Location20internal_constructor17h55a9de981ef3aeb7E
        local.get 4
        i32.const 36
        i32.add
        local.get 4
        i32.const 56
        i32.add
        i64.load
        i64.store align=4
        local.get 4
        local.get 2
        i32.store offset=24
        local.get 4
        i32.const 1048740
        i32.store offset=20
        local.get 4
        i32.const 1048740
        i32.store offset=16
        local.get 4
        local.get 4
        i64.load offset=48
        i64.store offset=28 align=4
        i32.const 0
        i32.load offset=1055716
        local.tee 3
        i32.const -1
        i32.le_s
        br_if 0 (;@2;)
        i32.const 0
        local.get 3
        i32.const 1
        i32.add
        local.tee 3
        i32.store offset=1055716
        block  ;; label = @3
          i32.const 0
          i32.load offset=1055724
          local.tee 2
          i32.eqz
          br_if 0 (;@3;)
          i32.const 0
          i32.load offset=1055720
          local.set 3
          local.get 4
          i32.const 8
          i32.add
          local.get 0
          local.get 1
          i32.load offset=16
          call_indirect (type 1)
          local.get 4
          local.get 4
          i64.load offset=8
          i64.store offset=16
          local.get 3
          local.get 4
          i32.const 16
          i32.add
          local.get 2
          i32.load offset=12
          call_indirect (type 1)
          i32.const 0
          i32.load offset=1055716
          local.set 3
        end
        i32.const 0
        local.get 3
        i32.const -1
        i32.add
        i32.store offset=1055716
        local.get 5
        i32.const 2
        i32.lt_u
        br_if 1 (;@1;)
      end
      unreachable
      unreachable
    end
    local.get 0
    local.get 1
    call $rust_panic
    unreachable)
  (func $_ZN89_$LT$std..panicking..continue_panic_fmt..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$9box_me_up17hf287a40b33981434E (type 1) (param i32 i32)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      local.get 1
      i32.load offset=4
      local.tee 3
      br_if 0 (;@1;)
      local.get 1
      i32.load
      local.set 3
      local.get 2
      i32.const 0
      i32.store offset=16
      local.get 2
      i64.const 1
      i64.store offset=8
      local.get 2
      local.get 2
      i32.const 8
      i32.add
      i32.store offset=20
      local.get 2
      i32.const 24
      i32.add
      i32.const 16
      i32.add
      local.get 3
      i32.const 16
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      i32.const 24
      i32.add
      i32.const 8
      i32.add
      local.tee 4
      local.get 3
      i32.const 8
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      local.get 3
      i64.load align=4
      i64.store offset=24
      local.get 2
      i32.const 20
      i32.add
      i32.const 1048576
      local.get 2
      i32.const 24
      i32.add
      call $_ZN4core3fmt5write17h8cfd01c67a4a46c9E
      drop
      local.get 4
      local.get 2
      i32.load offset=16
      i32.store
      local.get 2
      local.get 2
      i64.load offset=8
      i64.store offset=24
      block  ;; label = @2
        local.get 1
        i32.const 4
        i32.add
        local.tee 3
        i32.load
        local.tee 5
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        i32.const 8
        i32.add
        i32.load
        local.tee 6
        i32.eqz
        br_if 0 (;@2;)
        local.get 5
        local.get 6
        i32.const 1
        call $__rust_dealloc
      end
      local.get 3
      local.get 2
      i64.load offset=24
      i64.store align=4
      local.get 3
      i32.const 8
      i32.add
      local.get 4
      i32.load
      i32.store
      local.get 3
      i32.load
      local.set 3
    end
    local.get 1
    i32.const 1
    i32.store offset=4
    local.get 1
    i32.const 12
    i32.add
    i32.load
    local.set 4
    local.get 1
    i32.const 8
    i32.add
    local.tee 1
    i32.load
    local.set 5
    local.get 1
    i64.const 0
    i64.store align=4
    block  ;; label = @1
      i32.const 12
      i32.const 4
      call $__rust_alloc
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 4
      i32.store offset=8
      local.get 1
      local.get 5
      i32.store offset=4
      local.get 1
      local.get 3
      i32.store
      local.get 0
      i32.const 1049868
      i32.store offset=4
      local.get 0
      local.get 1
      i32.store
      local.get 2
      i32.const 48
      i32.add
      global.set 0
      return
    end
    i32.const 12
    i32.const 4
    call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
    unreachable)
  (func $_ZN89_$LT$std..panicking..continue_panic_fmt..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$3get17h0c4a906bac5cdb67E (type 1) (param i32 i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 1
    i32.const 4
    i32.add
    local.set 3
    block  ;; label = @1
      local.get 1
      i32.load offset=4
      br_if 0 (;@1;)
      local.get 1
      i32.load
      local.set 4
      local.get 2
      i32.const 0
      i32.store offset=16
      local.get 2
      i64.const 1
      i64.store offset=8
      local.get 2
      local.get 2
      i32.const 8
      i32.add
      i32.store offset=20
      local.get 2
      i32.const 24
      i32.add
      i32.const 16
      i32.add
      local.get 4
      i32.const 16
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      i32.const 24
      i32.add
      i32.const 8
      i32.add
      local.tee 5
      local.get 4
      i32.const 8
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      local.get 4
      i64.load align=4
      i64.store offset=24
      local.get 2
      i32.const 20
      i32.add
      i32.const 1048576
      local.get 2
      i32.const 24
      i32.add
      call $_ZN4core3fmt5write17h8cfd01c67a4a46c9E
      drop
      local.get 5
      local.get 2
      i32.load offset=16
      i32.store
      local.get 2
      local.get 2
      i64.load offset=8
      i64.store offset=24
      block  ;; label = @2
        local.get 3
        i32.load
        local.tee 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        i32.const 8
        i32.add
        i32.load
        local.tee 1
        i32.eqz
        br_if 0 (;@2;)
        local.get 4
        local.get 1
        i32.const 1
        call $__rust_dealloc
      end
      local.get 3
      local.get 2
      i64.load offset=24
      i64.store align=4
      local.get 3
      i32.const 8
      i32.add
      local.get 5
      i32.load
      i32.store
    end
    local.get 0
    i32.const 1049868
    i32.store offset=4
    local.get 0
    local.get 3
    i32.store
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func $_ZN91_$LT$std..panicking..begin_panic..PanicPayload$LT$A$GT$$u20$as$u20$core..panic..BoxMeUp$GT$9box_me_up17h76250b59498caedbE (type 1) (param i32 i32)
    (local i32 i32)
    local.get 1
    i32.load
    local.set 2
    local.get 1
    i32.const 0
    i32.store
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          i32.eqz
          br_if 0 (;@3;)
          local.get 1
          i32.load offset=4
          local.set 3
          i32.const 8
          i32.const 4
          call $__rust_alloc
          local.tee 1
          i32.eqz
          br_if 2 (;@1;)
          local.get 1
          local.get 3
          i32.store offset=4
          local.get 1
          local.get 2
          i32.store
          i32.const 1049904
          local.set 2
          br 1 (;@2;)
        end
        i32.const 1
        local.set 1
        i32.const 1049920
        local.set 2
      end
      local.get 0
      local.get 2
      i32.store offset=4
      local.get 0
      local.get 1
      i32.store
      return
    end
    i32.const 8
    i32.const 4
    call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
    unreachable)
  (func $_ZN91_$LT$std..panicking..begin_panic..PanicPayload$LT$A$GT$$u20$as$u20$core..panic..BoxMeUp$GT$3get17h634a6ff81d071bd5E (type 1) (param i32 i32)
    (local i32)
    local.get 0
    i32.const 1049904
    i32.const 1049920
    local.get 1
    i32.load
    local.tee 2
    select
    i32.store offset=4
    local.get 0
    local.get 1
    i32.const 1048740
    local.get 2
    select
    i32.store)
  (func $rust_panic (type 1) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.store offset=12
    local.get 2
    local.get 0
    i32.store offset=8
    local.get 2
    i32.const 8
    i32.add
    call $__rust_start_panic
    drop
    unreachable
    unreachable)
  (func $_ZN3std2rt19lang_start_internal17he5cddafed9dd65a0E (type 10) (param i32 i32 i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 4
    global.set 0
    local.get 4
    local.get 1
    i32.store offset=4
    local.get 4
    local.get 0
    i32.store
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            i32.const 4
            i32.const 1
            call $__rust_alloc
            local.tee 0
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            i32.const 1852399981
            i32.store align=1
            local.get 4
            i64.const 17179869188
            i64.store offset=12 align=4
            local.get 4
            local.get 0
            i32.store offset=8
            local.get 4
            i32.const 8
            i32.add
            call $_ZN3std6thread6Thread3new17h0a174627a8dd7c63E
            local.set 1
            i32.const 0
            local.set 0
            block  ;; label = @5
              block  ;; label = @6
                i32.const 0
                i32.load offset=1055728
                i32.const 1
                i32.ne
                br_if 0 (;@6;)
                i32.const 0
                i32.load offset=1055732
                local.set 0
                br 1 (;@5;)
              end
              i32.const 0
              i64.const 1
              i64.store offset=1055728 align=4
              i32.const 0
              i32.const 0
              i32.store offset=1055736
            end
            block  ;; label = @5
              local.get 0
              i32.const 0
              i32.lt_s
              br_if 0 (;@5;)
              local.get 0
              i32.const 2147483647
              i32.eq
              br_if 0 (;@5;)
              i32.const 0
              i32.load offset=1055736
              br_if 2 (;@3;)
              local.get 0
              br_if 3 (;@2;)
              i32.const 0
              local.set 0
              i32.const 0
              local.get 1
              i32.store offset=1055736
              i32.const 0
              i32.const 0
              i32.store offset=1055732
              local.get 4
              i32.const 0
              i32.store offset=24
              local.get 4
              i32.const 0
              i32.store offset=28
              local.get 4
              local.get 4
              i32.store offset=8
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      i32.const 9
                      local.get 4
                      i32.const 8
                      i32.add
                      local.get 4
                      i32.const 24
                      i32.add
                      local.get 4
                      i32.const 28
                      i32.add
                      call $__rust_maybe_catch_panic
                      i32.eqz
                      br_if 0 (;@9;)
                      i32.const 0
                      i32.load offset=1056192
                      i32.const 1
                      i32.ne
                      br_if 1 (;@8;)
                      i32.const 0
                      i32.load offset=1056196
                      i32.const -1
                      i32.add
                      local.set 0
                      br 2 (;@7;)
                    end
                    local.get 4
                    i32.load offset=8
                    local.set 1
                    br 2 (;@6;)
                  end
                  i32.const 0
                  i64.const 1
                  i64.store offset=1056192
                  i32.const -1
                  local.set 0
                end
                i32.const 0
                local.get 0
                i32.store offset=1056196
                i32.const 1
                local.set 0
                local.get 4
                i32.load offset=28
                local.set 5
                local.get 4
                i32.load offset=24
                local.set 1
              end
              i32.const 0
              i32.load offset=1055708
              i32.const 3
              i32.eq
              br_if 4 (;@1;)
              local.get 4
              i32.const 1
              i32.store8 offset=28
              local.get 4
              local.get 4
              i32.const 28
              i32.add
              i32.store offset=8
              i32.const 1055708
              i32.const 0
              local.get 4
              i32.const 8
              i32.add
              i32.const 1049464
              call $_ZN3std4sync4once4Once10call_inner17h92989b0c916494d1E
              br 4 (;@1;)
            end
            call $_ZN4core6result13unwrap_failed17h6d179f84b957db36E
            unreachable
          end
          i32.const 4
          i32.const 1
          call $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E
          unreachable
        end
        i32.const 1049808
        i32.const 38
        i32.const 1049792
        call $_ZN3std9panicking11begin_panic17h445db074890b5b34E
        unreachable
      end
      call $_ZN4core6result13unwrap_failed17h1984afaa28d103e8E
      unreachable
    end
    i32.const 101
    local.get 1
    local.get 0
    select
    local.set 6
    block  ;; label = @1
      local.get 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 5
      i32.load
      call_indirect (type 0)
      local.get 5
      i32.load offset=4
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 0
      local.get 5
      i32.load offset=8
      call $__rust_dealloc
    end
    local.get 4
    i32.const 32
    i32.add
    global.set 0
    local.get 6)
  (func $_ZN3std3sys4wasm7process8ExitCode6as_i3217h8c3f6c9d0d4bbc02E (type 2) (param i32) (result i32)
    local.get 0
    i32.load8_u)
  (func $__rust_maybe_catch_panic (type 10) (param i32 i32 i32 i32) (result i32)
    local.get 1
    local.get 0
    call_indirect (type 0)
    i32.const 0)
  (func $__rust_start_panic (type 2) (param i32) (result i32)
    unreachable
    unreachable)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc16malloc_alignment17h1f8103281a1655e6E (type 2) (param i32) (result i32)
    i32.const 8)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc6malloc17h7ef7d8a98d4afe8dE (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i64)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    block  ;; label = @17
                                      block  ;; label = @18
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              block  ;; label = @22
                                                block  ;; label = @23
                                                  block  ;; label = @24
                                                    block  ;; label = @25
                                                      block  ;; label = @26
                                                        block  ;; label = @27
                                                          block  ;; label = @28
                                                            block  ;; label = @29
                                                              block  ;; label = @30
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    block  ;; label = @33
                                                                      block  ;; label = @34
                                                                        block  ;; label = @35
                                                                          block  ;; label = @36
                                                                            block  ;; label = @37
                                                                              local.get 1
                                                                              i32.const 244
                                                                              i32.gt_u
                                                                              br_if 0 (;@37;)
                                                                              local.get 0
                                                                              i32.load
                                                                              local.tee 2
                                                                              i32.const 16
                                                                              local.get 1
                                                                              i32.const 11
                                                                              i32.add
                                                                              i32.const -8
                                                                              i32.and
                                                                              local.get 1
                                                                              i32.const 11
                                                                              i32.lt_u
                                                                              select
                                                                              local.tee 3
                                                                              i32.const 3
                                                                              i32.shr_u
                                                                              local.tee 4
                                                                              i32.const 31
                                                                              i32.and
                                                                              local.tee 5
                                                                              i32.shr_u
                                                                              local.tee 1
                                                                              i32.const 3
                                                                              i32.and
                                                                              i32.eqz
                                                                              br_if 1 (;@36;)
                                                                              local.get 0
                                                                              local.get 1
                                                                              i32.const -1
                                                                              i32.xor
                                                                              i32.const 1
                                                                              i32.and
                                                                              local.get 4
                                                                              i32.add
                                                                              local.tee 3
                                                                              i32.const 3
                                                                              i32.shl
                                                                              i32.add
                                                                              local.tee 6
                                                                              i32.const 16
                                                                              i32.add
                                                                              i32.load
                                                                              local.tee 1
                                                                              i32.const 8
                                                                              i32.add
                                                                              local.set 4
                                                                              local.get 1
                                                                              i32.load offset=8
                                                                              local.tee 5
                                                                              local.get 6
                                                                              i32.const 8
                                                                              i32.add
                                                                              local.tee 6
                                                                              i32.eq
                                                                              br_if 2 (;@35;)
                                                                              local.get 5
                                                                              local.get 6
                                                                              i32.store offset=12
                                                                              local.get 6
                                                                              i32.const 8
                                                                              i32.add
                                                                              local.get 5
                                                                              i32.store
                                                                              br 3 (;@34;)
                                                                            end
                                                                            i32.const 0
                                                                            local.set 4
                                                                            local.get 1
                                                                            i32.const -65588
                                                                            i32.gt_u
                                                                            br_if 35 (;@1;)
                                                                            local.get 1
                                                                            i32.const 11
                                                                            i32.add
                                                                            local.tee 1
                                                                            i32.const -8
                                                                            i32.and
                                                                            local.set 3
                                                                            local.get 0
                                                                            i32.load offset=4
                                                                            local.tee 7
                                                                            i32.eqz
                                                                            br_if 9 (;@27;)
                                                                            i32.const 0
                                                                            local.set 8
                                                                            block  ;; label = @37
                                                                              local.get 1
                                                                              i32.const 8
                                                                              i32.shr_u
                                                                              local.tee 1
                                                                              i32.eqz
                                                                              br_if 0 (;@37;)
                                                                              i32.const 31
                                                                              local.set 8
                                                                              local.get 3
                                                                              i32.const 16777215
                                                                              i32.gt_u
                                                                              br_if 0 (;@37;)
                                                                              local.get 3
                                                                              i32.const 38
                                                                              local.get 1
                                                                              i32.clz
                                                                              local.tee 1
                                                                              i32.sub
                                                                              i32.const 31
                                                                              i32.and
                                                                              i32.shr_u
                                                                              i32.const 1
                                                                              i32.and
                                                                              i32.const 31
                                                                              local.get 1
                                                                              i32.sub
                                                                              i32.const 1
                                                                              i32.shl
                                                                              i32.or
                                                                              local.set 8
                                                                            end
                                                                            i32.const 0
                                                                            local.get 3
                                                                            i32.sub
                                                                            local.set 4
                                                                            local.get 0
                                                                            local.get 8
                                                                            i32.const 2
                                                                            i32.shl
                                                                            i32.add
                                                                            i32.const 272
                                                                            i32.add
                                                                            i32.load
                                                                            local.tee 1
                                                                            i32.eqz
                                                                            br_if 6 (;@30;)
                                                                            i32.const 0
                                                                            local.set 5
                                                                            local.get 3
                                                                            i32.const 0
                                                                            i32.const 25
                                                                            local.get 8
                                                                            i32.const 1
                                                                            i32.shr_u
                                                                            i32.sub
                                                                            i32.const 31
                                                                            i32.and
                                                                            local.get 8
                                                                            i32.const 31
                                                                            i32.eq
                                                                            select
                                                                            i32.shl
                                                                            local.set 6
                                                                            i32.const 0
                                                                            local.set 2
                                                                            loop  ;; label = @37
                                                                              block  ;; label = @38
                                                                                local.get 1
                                                                                i32.load offset=4
                                                                                i32.const -8
                                                                                i32.and
                                                                                local.tee 9
                                                                                local.get 3
                                                                                i32.lt_u
                                                                                br_if 0 (;@38;)
                                                                                local.get 9
                                                                                local.get 3
                                                                                i32.sub
                                                                                local.tee 9
                                                                                local.get 4
                                                                                i32.ge_u
                                                                                br_if 0 (;@38;)
                                                                                local.get 9
                                                                                local.set 4
                                                                                local.get 1
                                                                                local.set 2
                                                                                local.get 9
                                                                                i32.eqz
                                                                                br_if 6 (;@32;)
                                                                              end
                                                                              local.get 1
                                                                              i32.const 20
                                                                              i32.add
                                                                              i32.load
                                                                              local.tee 9
                                                                              local.get 5
                                                                              local.get 9
                                                                              local.get 1
                                                                              local.get 6
                                                                              i32.const 29
                                                                              i32.shr_u
                                                                              i32.const 4
                                                                              i32.and
                                                                              i32.add
                                                                              i32.const 16
                                                                              i32.add
                                                                              i32.load
                                                                              local.tee 1
                                                                              i32.ne
                                                                              select
                                                                              local.get 5
                                                                              local.get 9
                                                                              select
                                                                              local.set 5
                                                                              local.get 6
                                                                              i32.const 1
                                                                              i32.shl
                                                                              local.set 6
                                                                              local.get 1
                                                                              br_if 0 (;@37;)
                                                                            end
                                                                            local.get 5
                                                                            i32.eqz
                                                                            br_if 5 (;@31;)
                                                                            local.get 5
                                                                            local.set 1
                                                                            br 7 (;@29;)
                                                                          end
                                                                          local.get 3
                                                                          local.get 0
                                                                          i32.load offset=400
                                                                          i32.le_u
                                                                          br_if 8 (;@27;)
                                                                          local.get 1
                                                                          i32.eqz
                                                                          br_if 2 (;@33;)
                                                                          local.get 0
                                                                          local.get 1
                                                                          local.get 5
                                                                          i32.shl
                                                                          i32.const 2
                                                                          local.get 5
                                                                          i32.shl
                                                                          local.tee 1
                                                                          i32.const 0
                                                                          local.get 1
                                                                          i32.sub
                                                                          i32.or
                                                                          i32.and
                                                                          local.tee 1
                                                                          i32.const 0
                                                                          local.get 1
                                                                          i32.sub
                                                                          i32.and
                                                                          i32.ctz
                                                                          local.tee 4
                                                                          i32.const 3
                                                                          i32.shl
                                                                          i32.add
                                                                          local.tee 6
                                                                          i32.const 16
                                                                          i32.add
                                                                          i32.load
                                                                          local.tee 1
                                                                          i32.load offset=8
                                                                          local.tee 5
                                                                          local.get 6
                                                                          i32.const 8
                                                                          i32.add
                                                                          local.tee 6
                                                                          i32.eq
                                                                          br_if 10 (;@25;)
                                                                          local.get 5
                                                                          local.get 6
                                                                          i32.store offset=12
                                                                          local.get 6
                                                                          i32.const 8
                                                                          i32.add
                                                                          local.get 5
                                                                          i32.store
                                                                          br 11 (;@24;)
                                                                        end
                                                                        local.get 0
                                                                        local.get 2
                                                                        i32.const -2
                                                                        local.get 3
                                                                        i32.rotl
                                                                        i32.and
                                                                        i32.store
                                                                      end
                                                                      local.get 1
                                                                      local.get 3
                                                                      i32.const 3
                                                                      i32.shl
                                                                      local.tee 3
                                                                      i32.const 3
                                                                      i32.or
                                                                      i32.store offset=4
                                                                      local.get 1
                                                                      local.get 3
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 1
                                                                      i32.load offset=4
                                                                      i32.const 1
                                                                      i32.or
                                                                      i32.store offset=4
                                                                      br 32 (;@1;)
                                                                    end
                                                                    local.get 0
                                                                    i32.load offset=4
                                                                    local.tee 1
                                                                    i32.eqz
                                                                    br_if 5 (;@27;)
                                                                    local.get 0
                                                                    local.get 1
                                                                    i32.const 0
                                                                    local.get 1
                                                                    i32.sub
                                                                    i32.and
                                                                    i32.ctz
                                                                    i32.const 2
                                                                    i32.shl
                                                                    i32.add
                                                                    i32.const 272
                                                                    i32.add
                                                                    i32.load
                                                                    local.tee 2
                                                                    i32.load offset=4
                                                                    i32.const -8
                                                                    i32.and
                                                                    local.get 3
                                                                    i32.sub
                                                                    local.set 5
                                                                    local.get 2
                                                                    local.set 6
                                                                    local.get 2
                                                                    i32.load offset=16
                                                                    local.tee 1
                                                                    i32.eqz
                                                                    br_if 20 (;@12;)
                                                                    i32.const 1
                                                                    local.set 4
                                                                    br 21 (;@11;)
                                                                  end
                                                                  i32.const 0
                                                                  local.set 4
                                                                  local.get 1
                                                                  local.set 2
                                                                  br 2 (;@29;)
                                                                end
                                                                local.get 2
                                                                br_if 2 (;@28;)
                                                              end
                                                              i32.const 0
                                                              local.set 2
                                                              i32.const 2
                                                              local.get 8
                                                              i32.const 31
                                                              i32.and
                                                              i32.shl
                                                              local.tee 1
                                                              i32.const 0
                                                              local.get 1
                                                              i32.sub
                                                              i32.or
                                                              local.get 7
                                                              i32.and
                                                              local.tee 1
                                                              i32.eqz
                                                              br_if 2 (;@27;)
                                                              local.get 0
                                                              local.get 1
                                                              i32.const 0
                                                              local.get 1
                                                              i32.sub
                                                              i32.and
                                                              i32.ctz
                                                              i32.const 2
                                                              i32.shl
                                                              i32.add
                                                              i32.const 272
                                                              i32.add
                                                              i32.load
                                                              local.tee 1
                                                              i32.eqz
                                                              br_if 2 (;@27;)
                                                            end
                                                            loop  ;; label = @29
                                                              local.get 1
                                                              i32.load offset=4
                                                              i32.const -8
                                                              i32.and
                                                              local.tee 5
                                                              local.get 3
                                                              i32.ge_u
                                                              local.get 5
                                                              local.get 3
                                                              i32.sub
                                                              local.tee 9
                                                              local.get 4
                                                              i32.lt_u
                                                              i32.and
                                                              local.set 6
                                                              block  ;; label = @30
                                                                local.get 1
                                                                i32.load offset=16
                                                                local.tee 5
                                                                br_if 0 (;@30;)
                                                                local.get 1
                                                                i32.const 20
                                                                i32.add
                                                                i32.load
                                                                local.set 5
                                                              end
                                                              local.get 1
                                                              local.get 2
                                                              local.get 6
                                                              select
                                                              local.set 2
                                                              local.get 9
                                                              local.get 4
                                                              local.get 6
                                                              select
                                                              local.set 4
                                                              local.get 5
                                                              local.set 1
                                                              local.get 5
                                                              br_if 0 (;@29;)
                                                            end
                                                            local.get 2
                                                            i32.eqz
                                                            br_if 1 (;@27;)
                                                          end
                                                          local.get 0
                                                          i32.load offset=400
                                                          local.tee 1
                                                          local.get 3
                                                          i32.lt_u
                                                          br_if 1 (;@26;)
                                                          local.get 4
                                                          local.get 1
                                                          local.get 3
                                                          i32.sub
                                                          i32.lt_u
                                                          br_if 1 (;@26;)
                                                        end
                                                        block  ;; label = @27
                                                          block  ;; label = @28
                                                            block  ;; label = @29
                                                              block  ;; label = @30
                                                                local.get 0
                                                                i32.load offset=400
                                                                local.tee 4
                                                                local.get 3
                                                                i32.ge_u
                                                                br_if 0 (;@30;)
                                                                local.get 0
                                                                i32.load offset=404
                                                                local.tee 1
                                                                local.get 3
                                                                i32.le_u
                                                                br_if 1 (;@29;)
                                                                local.get 0
                                                                i32.const 404
                                                                i32.add
                                                                local.get 1
                                                                local.get 3
                                                                i32.sub
                                                                local.tee 4
                                                                i32.store
                                                                local.get 0
                                                                local.get 0
                                                                i32.load offset=412
                                                                local.tee 1
                                                                local.get 3
                                                                i32.add
                                                                local.tee 5
                                                                i32.store offset=412
                                                                local.get 5
                                                                local.get 4
                                                                i32.const 1
                                                                i32.or
                                                                i32.store offset=4
                                                                local.get 1
                                                                local.get 3
                                                                i32.const 3
                                                                i32.or
                                                                i32.store offset=4
                                                                local.get 1
                                                                i32.const 8
                                                                i32.add
                                                                return
                                                              end
                                                              local.get 0
                                                              i32.load offset=408
                                                              local.set 1
                                                              local.get 4
                                                              local.get 3
                                                              i32.sub
                                                              local.tee 5
                                                              i32.const 16
                                                              i32.ge_u
                                                              br_if 1 (;@28;)
                                                              local.get 0
                                                              i32.const 408
                                                              i32.add
                                                              i32.const 0
                                                              i32.store
                                                              local.get 0
                                                              i32.const 400
                                                              i32.add
                                                              i32.const 0
                                                              i32.store
                                                              local.get 1
                                                              local.get 4
                                                              i32.const 3
                                                              i32.or
                                                              i32.store offset=4
                                                              local.get 1
                                                              local.get 4
                                                              i32.add
                                                              local.tee 4
                                                              i32.const 4
                                                              i32.add
                                                              local.set 3
                                                              local.get 4
                                                              i32.load offset=4
                                                              i32.const 1
                                                              i32.or
                                                              local.set 4
                                                              br 2 (;@27;)
                                                            end
                                                            i32.const 0
                                                            local.set 4
                                                            local.get 3
                                                            i32.const 65583
                                                            i32.add
                                                            local.tee 5
                                                            i32.const 16
                                                            i32.shr_u
                                                            memory.grow
                                                            local.tee 1
                                                            i32.const -1
                                                            i32.eq
                                                            br_if 27 (;@1;)
                                                            local.get 1
                                                            i32.const 16
                                                            i32.shl
                                                            local.tee 2
                                                            i32.eqz
                                                            br_if 27 (;@1;)
                                                            local.get 0
                                                            local.get 0
                                                            i32.load offset=416
                                                            local.get 5
                                                            i32.const -65536
                                                            i32.and
                                                            local.tee 8
                                                            i32.add
                                                            local.tee 1
                                                            i32.store offset=416
                                                            local.get 0
                                                            local.get 0
                                                            i32.load offset=420
                                                            local.tee 5
                                                            local.get 1
                                                            local.get 1
                                                            local.get 5
                                                            i32.lt_u
                                                            select
                                                            i32.store offset=420
                                                            local.get 0
                                                            i32.load offset=412
                                                            local.tee 5
                                                            i32.eqz
                                                            br_if 9 (;@19;)
                                                            local.get 0
                                                            i32.const 424
                                                            i32.add
                                                            local.tee 7
                                                            local.set 1
                                                            loop  ;; label = @29
                                                              local.get 1
                                                              i32.load
                                                              local.tee 6
                                                              local.get 1
                                                              i32.load offset=4
                                                              local.tee 9
                                                              i32.add
                                                              local.get 2
                                                              i32.eq
                                                              br_if 11 (;@18;)
                                                              local.get 1
                                                              i32.load offset=8
                                                              local.tee 1
                                                              br_if 0 (;@29;)
                                                              br 19 (;@10;)
                                                            end
                                                          end
                                                          local.get 0
                                                          i32.const 400
                                                          i32.add
                                                          local.get 5
                                                          i32.store
                                                          local.get 0
                                                          i32.const 408
                                                          i32.add
                                                          local.get 1
                                                          local.get 3
                                                          i32.add
                                                          local.tee 6
                                                          i32.store
                                                          local.get 6
                                                          local.get 5
                                                          i32.const 1
                                                          i32.or
                                                          i32.store offset=4
                                                          local.get 1
                                                          local.get 4
                                                          i32.add
                                                          local.get 5
                                                          i32.store
                                                          local.get 3
                                                          i32.const 3
                                                          i32.or
                                                          local.set 4
                                                          local.get 1
                                                          i32.const 4
                                                          i32.add
                                                          local.set 3
                                                        end
                                                        local.get 3
                                                        local.get 4
                                                        i32.store
                                                        local.get 1
                                                        i32.const 8
                                                        i32.add
                                                        return
                                                      end
                                                      local.get 0
                                                      local.get 2
                                                      call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
                                                      local.get 4
                                                      i32.const 15
                                                      i32.gt_u
                                                      br_if 2 (;@23;)
                                                      local.get 2
                                                      local.get 4
                                                      local.get 3
                                                      i32.add
                                                      local.tee 1
                                                      i32.const 3
                                                      i32.or
                                                      i32.store offset=4
                                                      local.get 2
                                                      local.get 1
                                                      i32.add
                                                      local.tee 1
                                                      local.get 1
                                                      i32.load offset=4
                                                      i32.const 1
                                                      i32.or
                                                      i32.store offset=4
                                                      br 12 (;@13;)
                                                    end
                                                    local.get 0
                                                    local.get 2
                                                    i32.const -2
                                                    local.get 4
                                                    i32.rotl
                                                    i32.and
                                                    i32.store
                                                  end
                                                  local.get 1
                                                  i32.const 8
                                                  i32.add
                                                  local.set 5
                                                  local.get 1
                                                  local.get 3
                                                  i32.const 3
                                                  i32.or
                                                  i32.store offset=4
                                                  local.get 1
                                                  local.get 3
                                                  i32.add
                                                  local.tee 6
                                                  local.get 4
                                                  i32.const 3
                                                  i32.shl
                                                  local.tee 4
                                                  local.get 3
                                                  i32.sub
                                                  local.tee 3
                                                  i32.const 1
                                                  i32.or
                                                  i32.store offset=4
                                                  local.get 1
                                                  local.get 4
                                                  i32.add
                                                  local.get 3
                                                  i32.store
                                                  local.get 0
                                                  i32.const 400
                                                  i32.add
                                                  local.tee 2
                                                  i32.load
                                                  local.tee 1
                                                  i32.eqz
                                                  br_if 3 (;@20;)
                                                  local.get 0
                                                  local.get 1
                                                  i32.const 3
                                                  i32.shr_u
                                                  local.tee 9
                                                  i32.const 3
                                                  i32.shl
                                                  i32.add
                                                  i32.const 8
                                                  i32.add
                                                  local.set 4
                                                  local.get 0
                                                  i32.load offset=408
                                                  local.set 1
                                                  local.get 0
                                                  i32.load
                                                  local.tee 8
                                                  i32.const 1
                                                  local.get 9
                                                  i32.const 31
                                                  i32.and
                                                  i32.shl
                                                  local.tee 9
                                                  i32.and
                                                  i32.eqz
                                                  br_if 1 (;@22;)
                                                  local.get 4
                                                  i32.load offset=8
                                                  local.set 9
                                                  br 2 (;@21;)
                                                end
                                                local.get 2
                                                local.get 3
                                                i32.const 3
                                                i32.or
                                                i32.store offset=4
                                                local.get 2
                                                local.get 3
                                                i32.add
                                                local.tee 1
                                                local.get 4
                                                i32.const 1
                                                i32.or
                                                i32.store offset=4
                                                local.get 1
                                                local.get 4
                                                i32.add
                                                local.get 4
                                                i32.store
                                                local.get 4
                                                i32.const 255
                                                i32.gt_u
                                                br_if 5 (;@17;)
                                                local.get 0
                                                local.get 4
                                                i32.const 3
                                                i32.shr_u
                                                local.tee 4
                                                i32.const 3
                                                i32.shl
                                                i32.add
                                                i32.const 8
                                                i32.add
                                                local.set 3
                                                local.get 0
                                                i32.load
                                                local.tee 5
                                                i32.const 1
                                                local.get 4
                                                i32.const 31
                                                i32.and
                                                i32.shl
                                                local.tee 4
                                                i32.and
                                                i32.eqz
                                                br_if 7 (;@15;)
                                                local.get 3
                                                i32.load offset=8
                                                local.set 4
                                                br 8 (;@14;)
                                              end
                                              local.get 0
                                              local.get 8
                                              local.get 9
                                              i32.or
                                              i32.store
                                              local.get 4
                                              local.set 9
                                            end
                                            local.get 4
                                            local.get 1
                                            i32.store offset=8
                                            local.get 9
                                            local.get 1
                                            i32.store offset=12
                                            local.get 1
                                            local.get 4
                                            i32.store offset=12
                                            local.get 1
                                            local.get 9
                                            i32.store offset=8
                                          end
                                          local.get 0
                                          local.get 6
                                          i32.store offset=408
                                          local.get 2
                                          local.get 3
                                          i32.store
                                          local.get 5
                                          return
                                        end
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            local.get 0
                                            i32.load offset=444
                                            local.tee 1
                                            i32.eqz
                                            br_if 0 (;@20;)
                                            local.get 1
                                            local.get 2
                                            i32.le_u
                                            br_if 1 (;@19;)
                                          end
                                          local.get 0
                                          i32.const 444
                                          i32.add
                                          local.get 2
                                          i32.store
                                        end
                                        local.get 0
                                        i32.const 4095
                                        i32.store offset=448
                                        local.get 0
                                        local.get 2
                                        i32.store offset=424
                                        i32.const 0
                                        local.set 1
                                        local.get 0
                                        i32.const 436
                                        i32.add
                                        i32.const 0
                                        i32.store
                                        local.get 0
                                        i32.const 428
                                        i32.add
                                        local.get 8
                                        i32.store
                                        loop  ;; label = @19
                                          local.get 0
                                          local.get 1
                                          i32.add
                                          local.tee 5
                                          i32.const 16
                                          i32.add
                                          local.get 5
                                          i32.const 8
                                          i32.add
                                          local.tee 6
                                          i32.store
                                          local.get 5
                                          i32.const 20
                                          i32.add
                                          local.get 6
                                          i32.store
                                          local.get 1
                                          i32.const 8
                                          i32.add
                                          local.tee 1
                                          i32.const 256
                                          i32.ne
                                          br_if 0 (;@19;)
                                        end
                                        local.get 0
                                        i32.const 404
                                        i32.add
                                        local.get 8
                                        i32.const -40
                                        i32.add
                                        local.tee 1
                                        i32.store
                                        local.get 0
                                        i32.const 412
                                        i32.add
                                        local.get 2
                                        i32.store
                                        local.get 2
                                        local.get 1
                                        i32.const 1
                                        i32.or
                                        i32.store offset=4
                                        local.get 2
                                        local.get 1
                                        i32.add
                                        i32.const 40
                                        i32.store offset=4
                                        local.get 0
                                        i32.const 2097152
                                        i32.store offset=440
                                        br 9 (;@9;)
                                      end
                                      local.get 1
                                      i32.load offset=12
                                      i32.eqz
                                      br_if 1 (;@16;)
                                      br 7 (;@10;)
                                    end
                                    local.get 0
                                    local.get 1
                                    local.get 4
                                    call $_ZN8dlmalloc8dlmalloc8Dlmalloc18insert_large_chunk17he00600b9254d4992E
                                    br 3 (;@13;)
                                  end
                                  local.get 2
                                  local.get 5
                                  i32.le_u
                                  br_if 5 (;@10;)
                                  local.get 6
                                  local.get 5
                                  i32.gt_u
                                  br_if 5 (;@10;)
                                  local.get 1
                                  i32.const 4
                                  i32.add
                                  local.get 9
                                  local.get 8
                                  i32.add
                                  i32.store
                                  local.get 0
                                  i32.const 412
                                  i32.add
                                  local.tee 1
                                  local.get 1
                                  i32.load
                                  local.tee 1
                                  i32.const 15
                                  i32.add
                                  i32.const -8
                                  i32.and
                                  local.tee 5
                                  i32.const -8
                                  i32.add
                                  local.tee 6
                                  i32.store
                                  local.get 0
                                  i32.const 404
                                  i32.add
                                  local.tee 2
                                  local.get 2
                                  i32.load
                                  local.get 8
                                  i32.add
                                  local.tee 2
                                  local.get 1
                                  i32.const 8
                                  i32.add
                                  local.get 5
                                  i32.sub
                                  i32.add
                                  local.tee 5
                                  i32.store
                                  local.get 6
                                  local.get 5
                                  i32.const 1
                                  i32.or
                                  i32.store offset=4
                                  local.get 1
                                  local.get 2
                                  i32.add
                                  i32.const 40
                                  i32.store offset=4
                                  local.get 0
                                  i32.const 2097152
                                  i32.store offset=440
                                  br 6 (;@9;)
                                end
                                local.get 0
                                local.get 5
                                local.get 4
                                i32.or
                                i32.store
                                local.get 3
                                local.set 4
                              end
                              local.get 3
                              local.get 1
                              i32.store offset=8
                              local.get 4
                              local.get 1
                              i32.store offset=12
                              local.get 1
                              local.get 3
                              i32.store offset=12
                              local.get 1
                              local.get 4
                              i32.store offset=8
                            end
                            local.get 2
                            i32.const 8
                            i32.add
                            return
                          end
                          i32.const 0
                          local.set 4
                        end
                        loop  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  local.get 4
                                  br_table 1 (;@14;) 0 (;@15;) 0 (;@15;)
                                end
                                local.get 1
                                i32.load offset=4
                                i32.const -8
                                i32.and
                                local.get 3
                                i32.sub
                                local.tee 4
                                local.get 5
                                local.get 4
                                local.get 5
                                i32.lt_u
                                local.tee 4
                                select
                                local.set 5
                                local.get 1
                                local.get 6
                                local.get 4
                                select
                                local.set 6
                                local.get 1
                                local.tee 2
                                i32.load offset=16
                                local.tee 1
                                br_if 1 (;@13;)
                                i32.const 0
                                local.set 4
                                br 3 (;@11;)
                              end
                              local.get 2
                              i32.const 20
                              i32.add
                              i32.load
                              local.tee 1
                              br_if 1 (;@12;)
                              local.get 0
                              local.get 6
                              call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
                              block  ;; label = @14
                                block  ;; label = @15
                                  local.get 5
                                  i32.const 16
                                  i32.ge_u
                                  br_if 0 (;@15;)
                                  local.get 6
                                  local.get 5
                                  local.get 3
                                  i32.add
                                  local.tee 1
                                  i32.const 3
                                  i32.or
                                  i32.store offset=4
                                  local.get 6
                                  local.get 1
                                  i32.add
                                  local.tee 1
                                  local.get 1
                                  i32.load offset=4
                                  i32.const 1
                                  i32.or
                                  i32.store offset=4
                                  br 1 (;@14;)
                                end
                                local.get 6
                                local.get 3
                                i32.const 3
                                i32.or
                                i32.store offset=4
                                local.get 6
                                local.get 3
                                i32.add
                                local.tee 3
                                local.get 5
                                i32.const 1
                                i32.or
                                i32.store offset=4
                                local.get 3
                                local.get 5
                                i32.add
                                local.get 5
                                i32.store
                                block  ;; label = @15
                                  local.get 0
                                  i32.const 400
                                  i32.add
                                  local.tee 2
                                  i32.load
                                  local.tee 1
                                  i32.eqz
                                  br_if 0 (;@15;)
                                  local.get 0
                                  local.get 1
                                  i32.const 3
                                  i32.shr_u
                                  local.tee 9
                                  i32.const 3
                                  i32.shl
                                  i32.add
                                  i32.const 8
                                  i32.add
                                  local.set 4
                                  local.get 0
                                  i32.load offset=408
                                  local.set 1
                                  block  ;; label = @16
                                    block  ;; label = @17
                                      local.get 0
                                      i32.load
                                      local.tee 8
                                      i32.const 1
                                      local.get 9
                                      i32.const 31
                                      i32.and
                                      i32.shl
                                      local.tee 9
                                      i32.and
                                      i32.eqz
                                      br_if 0 (;@17;)
                                      local.get 4
                                      i32.load offset=8
                                      local.set 9
                                      br 1 (;@16;)
                                    end
                                    local.get 0
                                    local.get 8
                                    local.get 9
                                    i32.or
                                    i32.store
                                    local.get 4
                                    local.set 9
                                  end
                                  local.get 4
                                  local.get 1
                                  i32.store offset=8
                                  local.get 9
                                  local.get 1
                                  i32.store offset=12
                                  local.get 1
                                  local.get 4
                                  i32.store offset=12
                                  local.get 1
                                  local.get 9
                                  i32.store offset=8
                                end
                                local.get 0
                                local.get 3
                                i32.store offset=408
                                local.get 2
                                local.get 5
                                i32.store
                              end
                              local.get 6
                              i32.const 8
                              i32.add
                              return
                            end
                            i32.const 1
                            local.set 4
                            br 1 (;@11;)
                          end
                          i32.const 1
                          local.set 4
                          br 0 (;@11;)
                        end
                      end
                      local.get 0
                      local.get 0
                      i32.load offset=444
                      local.tee 1
                      local.get 2
                      local.get 1
                      local.get 2
                      i32.lt_u
                      select
                      i32.store offset=444
                      local.get 2
                      local.get 8
                      i32.add
                      local.set 9
                      local.get 7
                      local.set 6
                      block  ;; label = @10
                        block  ;; label = @11
                          loop  ;; label = @12
                            local.get 6
                            i32.load
                            local.get 9
                            i32.eq
                            br_if 1 (;@11;)
                            local.get 6
                            i32.load offset=8
                            local.tee 6
                            br_if 0 (;@12;)
                          end
                          local.get 7
                          local.set 1
                          br 1 (;@10;)
                        end
                        local.get 7
                        local.set 1
                        local.get 6
                        i32.load offset=12
                        br_if 0 (;@10;)
                        local.get 6
                        local.get 2
                        i32.store
                        local.get 6
                        local.get 6
                        i32.load offset=4
                        local.get 8
                        i32.add
                        i32.store offset=4
                        local.get 2
                        local.get 3
                        i32.const 3
                        i32.or
                        i32.store offset=4
                        local.get 2
                        local.get 3
                        i32.add
                        local.set 1
                        local.get 9
                        local.get 2
                        i32.sub
                        local.get 3
                        i32.sub
                        local.set 3
                        local.get 0
                        i32.const 412
                        i32.add
                        local.tee 4
                        i32.load
                        local.get 9
                        i32.eq
                        br_if 2 (;@8;)
                        local.get 0
                        i32.load offset=408
                        local.get 9
                        i32.eq
                        br_if 3 (;@7;)
                        local.get 9
                        i32.load offset=4
                        local.tee 4
                        i32.const 3
                        i32.and
                        i32.const 1
                        i32.ne
                        br_if 7 (;@3;)
                        local.get 4
                        i32.const -8
                        i32.and
                        local.tee 5
                        i32.const 255
                        i32.gt_u
                        br_if 4 (;@6;)
                        local.get 9
                        i32.load offset=12
                        local.tee 6
                        local.get 9
                        i32.load offset=8
                        local.tee 8
                        i32.eq
                        br_if 5 (;@5;)
                        local.get 8
                        local.get 6
                        i32.store offset=12
                        local.get 6
                        local.get 8
                        i32.store offset=8
                        br 6 (;@4;)
                      end
                      block  ;; label = @10
                        loop  ;; label = @11
                          block  ;; label = @12
                            local.get 1
                            i32.load
                            local.tee 6
                            local.get 5
                            i32.gt_u
                            br_if 0 (;@12;)
                            local.get 6
                            local.get 1
                            i32.load offset=4
                            i32.add
                            local.tee 6
                            local.get 5
                            i32.gt_u
                            br_if 2 (;@10;)
                          end
                          local.get 1
                          i32.load offset=8
                          local.set 1
                          br 0 (;@11;)
                        end
                      end
                      local.get 0
                      i32.const 404
                      i32.add
                      local.get 8
                      i32.const -40
                      i32.add
                      local.tee 1
                      i32.store
                      local.get 0
                      i32.const 412
                      i32.add
                      local.get 2
                      i32.store
                      local.get 2
                      local.get 1
                      i32.const 1
                      i32.or
                      i32.store offset=4
                      local.get 2
                      local.get 1
                      i32.add
                      i32.const 40
                      i32.store offset=4
                      local.get 0
                      i32.const 2097152
                      i32.store offset=440
                      local.get 5
                      local.get 6
                      i32.const -32
                      i32.add
                      i32.const -8
                      i32.and
                      i32.const -8
                      i32.add
                      local.tee 1
                      local.get 1
                      local.get 5
                      i32.const 16
                      i32.add
                      i32.lt_u
                      select
                      local.tee 9
                      i32.const 27
                      i32.store offset=4
                      local.get 7
                      i64.load align=4
                      local.set 10
                      local.get 9
                      i32.const 16
                      i32.add
                      local.get 7
                      i32.const 8
                      i32.add
                      i64.load align=4
                      i64.store align=4
                      local.get 9
                      local.get 10
                      i64.store offset=8 align=4
                      local.get 0
                      i32.const 436
                      i32.add
                      i32.const 0
                      i32.store
                      local.get 0
                      i32.const 428
                      i32.add
                      local.get 8
                      i32.store
                      local.get 0
                      i32.const 424
                      i32.add
                      local.get 2
                      i32.store
                      local.get 0
                      i32.const 432
                      i32.add
                      local.get 9
                      i32.const 8
                      i32.add
                      i32.store
                      local.get 9
                      i32.const 28
                      i32.add
                      local.set 1
                      loop  ;; label = @10
                        local.get 1
                        i32.const 7
                        i32.store
                        local.get 6
                        local.get 1
                        i32.const 4
                        i32.add
                        local.tee 1
                        i32.gt_u
                        br_if 0 (;@10;)
                      end
                      local.get 9
                      local.get 5
                      i32.eq
                      br_if 0 (;@9;)
                      local.get 9
                      local.get 9
                      i32.load offset=4
                      i32.const -2
                      i32.and
                      i32.store offset=4
                      local.get 5
                      local.get 9
                      local.get 5
                      i32.sub
                      local.tee 1
                      i32.const 1
                      i32.or
                      i32.store offset=4
                      local.get 9
                      local.get 1
                      i32.store
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            local.get 1
                            i32.const 255
                            i32.gt_u
                            br_if 0 (;@12;)
                            local.get 0
                            local.get 1
                            i32.const 3
                            i32.shr_u
                            local.tee 6
                            i32.const 3
                            i32.shl
                            i32.add
                            i32.const 8
                            i32.add
                            local.set 1
                            local.get 0
                            i32.load
                            local.tee 2
                            i32.const 1
                            local.get 6
                            i32.const 31
                            i32.and
                            i32.shl
                            local.tee 6
                            i32.and
                            i32.eqz
                            br_if 1 (;@11;)
                            local.get 1
                            i32.load offset=8
                            local.set 6
                            br 2 (;@10;)
                          end
                          local.get 0
                          local.get 5
                          local.get 1
                          call $_ZN8dlmalloc8dlmalloc8Dlmalloc18insert_large_chunk17he00600b9254d4992E
                          br 2 (;@9;)
                        end
                        local.get 0
                        local.get 2
                        local.get 6
                        i32.or
                        i32.store
                        local.get 1
                        local.set 6
                      end
                      local.get 1
                      local.get 5
                      i32.store offset=8
                      local.get 6
                      local.get 5
                      i32.store offset=12
                      local.get 5
                      local.get 1
                      i32.store offset=12
                      local.get 5
                      local.get 6
                      i32.store offset=8
                    end
                    local.get 0
                    i32.const 404
                    i32.add
                    local.tee 1
                    i32.load
                    local.tee 5
                    local.get 3
                    i32.le_u
                    br_if 7 (;@1;)
                    local.get 1
                    local.get 5
                    local.get 3
                    i32.sub
                    local.tee 4
                    i32.store
                    local.get 0
                    i32.const 412
                    i32.add
                    local.tee 1
                    local.get 1
                    i32.load
                    local.tee 1
                    local.get 3
                    i32.add
                    local.tee 5
                    i32.store
                    local.get 5
                    local.get 4
                    i32.const 1
                    i32.or
                    i32.store offset=4
                    local.get 1
                    local.get 3
                    i32.const 3
                    i32.or
                    i32.store offset=4
                    local.get 1
                    i32.const 8
                    i32.add
                    return
                  end
                  local.get 4
                  local.get 1
                  i32.store
                  local.get 0
                  i32.const 404
                  i32.add
                  local.tee 4
                  local.get 4
                  i32.load
                  local.get 3
                  i32.add
                  local.tee 3
                  i32.store
                  local.get 1
                  local.get 3
                  i32.const 1
                  i32.or
                  i32.store offset=4
                  br 5 (;@2;)
                end
                local.get 0
                i32.const 408
                i32.add
                local.get 1
                i32.store
                local.get 0
                i32.const 400
                i32.add
                local.tee 4
                local.get 4
                i32.load
                local.get 3
                i32.add
                local.tee 3
                i32.store
                local.get 1
                local.get 3
                i32.const 1
                i32.or
                i32.store offset=4
                local.get 1
                local.get 3
                i32.add
                local.get 3
                i32.store
                br 4 (;@2;)
              end
              local.get 0
              local.get 9
              call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
              br 1 (;@4;)
            end
            local.get 0
            local.get 0
            i32.load
            i32.const -2
            local.get 4
            i32.const 3
            i32.shr_u
            i32.rotl
            i32.and
            i32.store
          end
          local.get 5
          local.get 3
          i32.add
          local.set 3
          local.get 9
          local.get 5
          i32.add
          local.set 9
        end
        local.get 9
        local.get 9
        i32.load offset=4
        i32.const -2
        i32.and
        i32.store offset=4
        local.get 1
        local.get 3
        i32.const 1
        i32.or
        i32.store offset=4
        local.get 1
        local.get 3
        i32.add
        local.get 3
        i32.store
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 3
              i32.const 255
              i32.gt_u
              br_if 0 (;@5;)
              local.get 0
              local.get 3
              i32.const 3
              i32.shr_u
              local.tee 4
              i32.const 3
              i32.shl
              i32.add
              i32.const 8
              i32.add
              local.set 3
              local.get 0
              i32.load
              local.tee 5
              i32.const 1
              local.get 4
              i32.const 31
              i32.and
              i32.shl
              local.tee 4
              i32.and
              i32.eqz
              br_if 1 (;@4;)
              local.get 3
              i32.load offset=8
              local.set 4
              br 2 (;@3;)
            end
            local.get 0
            local.get 1
            local.get 3
            call $_ZN8dlmalloc8dlmalloc8Dlmalloc18insert_large_chunk17he00600b9254d4992E
            br 2 (;@2;)
          end
          local.get 0
          local.get 5
          local.get 4
          i32.or
          i32.store
          local.get 3
          local.set 4
        end
        local.get 3
        local.get 1
        i32.store offset=8
        local.get 4
        local.get 1
        i32.store offset=12
        local.get 1
        local.get 3
        i32.store offset=12
        local.get 1
        local.get 4
        i32.store offset=8
      end
      local.get 2
      i32.const 8
      i32.add
      return
    end
    local.get 4)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E (type 1) (param i32 i32)
    (local i32 i32 i32 i32 i32)
    local.get 1
    i32.load offset=24
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.load offset=12
            local.tee 3
            local.get 1
            i32.eq
            br_if 0 (;@4;)
            local.get 1
            i32.load offset=8
            local.tee 4
            local.get 3
            i32.store offset=12
            local.get 3
            local.get 4
            i32.store offset=8
            local.get 2
            br_if 1 (;@3;)
            br 2 (;@2;)
          end
          block  ;; label = @4
            local.get 1
            i32.const 20
            i32.const 16
            local.get 1
            i32.const 20
            i32.add
            local.tee 3
            i32.load
            local.tee 5
            select
            i32.add
            i32.load
            local.tee 4
            i32.eqz
            br_if 0 (;@4;)
            local.get 3
            local.get 1
            i32.const 16
            i32.add
            local.get 5
            select
            local.set 5
            block  ;; label = @5
              loop  ;; label = @6
                local.get 5
                local.set 6
                block  ;; label = @7
                  local.get 4
                  local.tee 3
                  i32.const 20
                  i32.add
                  local.tee 5
                  i32.load
                  local.tee 4
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 4
                  br_if 1 (;@6;)
                  br 2 (;@5;)
                end
                local.get 3
                i32.const 16
                i32.add
                local.set 5
                local.get 3
                i32.load offset=16
                local.tee 4
                br_if 0 (;@6;)
              end
            end
            local.get 6
            i32.const 0
            i32.store
            local.get 2
            br_if 1 (;@3;)
            br 2 (;@2;)
          end
          i32.const 0
          local.set 3
          local.get 2
          i32.eqz
          br_if 1 (;@2;)
        end
        block  ;; label = @3
          block  ;; label = @4
            local.get 0
            local.get 1
            i32.load offset=28
            i32.const 2
            i32.shl
            i32.add
            i32.const 272
            i32.add
            local.tee 4
            i32.load
            local.get 1
            i32.eq
            br_if 0 (;@4;)
            local.get 2
            i32.const 16
            i32.const 20
            local.get 2
            i32.load offset=16
            local.get 1
            i32.eq
            select
            i32.add
            local.get 3
            i32.store
            local.get 3
            br_if 1 (;@3;)
            br 2 (;@2;)
          end
          local.get 4
          local.get 3
          i32.store
          local.get 3
          i32.eqz
          br_if 2 (;@1;)
        end
        local.get 3
        local.get 2
        i32.store offset=24
        block  ;; label = @3
          local.get 1
          i32.load offset=16
          local.tee 4
          i32.eqz
          br_if 0 (;@3;)
          local.get 3
          local.get 4
          i32.store offset=16
          local.get 4
          local.get 3
          i32.store offset=24
        end
        local.get 1
        i32.const 20
        i32.add
        i32.load
        local.tee 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 3
        i32.const 20
        i32.add
        local.get 4
        i32.store
        local.get 4
        local.get 3
        i32.store offset=24
      end
      return
    end
    local.get 0
    local.get 0
    i32.load offset=4
    i32.const -2
    local.get 1
    i32.const 28
    i32.add
    i32.load
    i32.rotl
    i32.and
    i32.store offset=4)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc18insert_large_chunk17he00600b9254d4992E (type 8) (param i32 i32 i32)
    (local i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 8
        i32.shr_u
        local.tee 3
        i32.eqz
        br_if 0 (;@2;)
        i32.const 31
        local.set 4
        local.get 2
        i32.const 16777215
        i32.gt_u
        br_if 1 (;@1;)
        local.get 2
        i32.const 38
        local.get 3
        i32.clz
        local.tee 4
        i32.sub
        i32.const 31
        i32.and
        i32.shr_u
        i32.const 1
        i32.and
        i32.const 31
        local.get 4
        i32.sub
        i32.const 1
        i32.shl
        i32.or
        local.set 4
        br 1 (;@1;)
      end
      i32.const 0
      local.set 4
    end
    local.get 1
    i64.const 0
    i64.store offset=16 align=4
    local.get 1
    local.get 4
    i32.store offset=28
    local.get 0
    local.get 4
    i32.const 2
    i32.shl
    i32.add
    i32.const 272
    i32.add
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.load offset=4
              local.tee 5
              i32.const 1
              local.get 4
              i32.const 31
              i32.and
              i32.shl
              local.tee 6
              i32.and
              i32.eqz
              br_if 0 (;@5;)
              local.get 3
              i32.load
              local.tee 3
              i32.load offset=4
              i32.const -8
              i32.and
              local.get 2
              i32.ne
              br_if 1 (;@4;)
              local.get 3
              local.set 4
              br 2 (;@3;)
            end
            local.get 0
            i32.const 4
            i32.add
            local.get 5
            local.get 6
            i32.or
            i32.store
            local.get 3
            local.get 1
            i32.store
            local.get 1
            local.get 3
            i32.store offset=24
            br 3 (;@1;)
          end
          local.get 2
          i32.const 0
          i32.const 25
          local.get 4
          i32.const 1
          i32.shr_u
          i32.sub
          i32.const 31
          i32.and
          local.get 4
          i32.const 31
          i32.eq
          select
          i32.shl
          local.set 0
          loop  ;; label = @4
            local.get 3
            local.get 0
            i32.const 29
            i32.shr_u
            i32.const 4
            i32.and
            i32.add
            i32.const 16
            i32.add
            local.tee 5
            i32.load
            local.tee 4
            i32.eqz
            br_if 2 (;@2;)
            local.get 0
            i32.const 1
            i32.shl
            local.set 0
            local.get 4
            local.set 3
            local.get 4
            i32.load offset=4
            i32.const -8
            i32.and
            local.get 2
            i32.ne
            br_if 0 (;@4;)
          end
        end
        local.get 4
        i32.load offset=8
        local.tee 0
        local.get 1
        i32.store offset=12
        local.get 4
        local.get 1
        i32.store offset=8
        local.get 1
        i32.const 0
        i32.store offset=24
        local.get 1
        local.get 4
        i32.store offset=12
        local.get 1
        local.get 0
        i32.store offset=8
        return
      end
      local.get 5
      local.get 1
      i32.store
      local.get 1
      local.get 3
      i32.store offset=24
    end
    local.get 1
    local.get 1
    i32.store offset=12
    local.get 1
    local.get 1
    i32.store offset=8)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc7realloc17h60cd22bbd1e08b5cE (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    i32.const 0
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.const -65587
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 16
              local.get 2
              i32.const 11
              i32.add
              i32.const -8
              i32.and
              local.get 2
              i32.const 11
              i32.lt_u
              select
              local.set 4
              local.get 1
              i32.const -4
              i32.add
              local.tee 5
              i32.load
              local.tee 6
              i32.const -8
              i32.and
              local.set 7
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 6
                        i32.const 3
                        i32.and
                        i32.eqz
                        br_if 0 (;@10;)
                        local.get 1
                        i32.const -8
                        i32.add
                        local.tee 8
                        local.get 7
                        i32.add
                        local.set 9
                        local.get 7
                        local.get 4
                        i32.ge_u
                        br_if 1 (;@9;)
                        local.get 0
                        i32.load offset=412
                        local.get 9
                        i32.eq
                        br_if 2 (;@8;)
                        local.get 0
                        i32.load offset=408
                        local.get 9
                        i32.eq
                        br_if 3 (;@7;)
                        local.get 9
                        i32.load offset=4
                        local.tee 6
                        i32.const 2
                        i32.and
                        br_if 4 (;@6;)
                        local.get 6
                        i32.const -8
                        i32.and
                        local.tee 10
                        local.get 7
                        i32.add
                        local.tee 7
                        local.get 4
                        i32.lt_u
                        br_if 4 (;@6;)
                        local.get 7
                        local.get 4
                        i32.sub
                        local.set 2
                        local.get 10
                        i32.const 255
                        i32.gt_u
                        br_if 7 (;@3;)
                        local.get 9
                        i32.load offset=12
                        local.tee 3
                        local.get 9
                        i32.load offset=8
                        local.tee 9
                        i32.eq
                        br_if 8 (;@2;)
                        local.get 9
                        local.get 3
                        i32.store offset=12
                        local.get 3
                        local.get 9
                        i32.store offset=8
                        br 9 (;@1;)
                      end
                      local.get 4
                      i32.const 256
                      i32.lt_u
                      br_if 3 (;@6;)
                      local.get 7
                      local.get 4
                      i32.const 4
                      i32.or
                      i32.lt_u
                      br_if 3 (;@6;)
                      local.get 7
                      local.get 4
                      i32.sub
                      i32.const 131073
                      i32.ge_u
                      br_if 3 (;@6;)
                      local.get 1
                      return
                    end
                    block  ;; label = @9
                      local.get 7
                      local.get 4
                      i32.sub
                      local.tee 2
                      i32.const 16
                      i32.ge_u
                      br_if 0 (;@9;)
                      local.get 1
                      return
                    end
                    local.get 5
                    local.get 4
                    local.get 6
                    i32.const 1
                    i32.and
                    i32.or
                    i32.const 2
                    i32.or
                    i32.store
                    local.get 8
                    local.get 4
                    i32.add
                    local.tee 3
                    local.get 2
                    i32.const 3
                    i32.or
                    i32.store offset=4
                    local.get 9
                    local.get 9
                    i32.load offset=4
                    i32.const 1
                    i32.or
                    i32.store offset=4
                    local.get 0
                    local.get 3
                    local.get 2
                    call $_ZN8dlmalloc8dlmalloc8Dlmalloc13dispose_chunk17h08c63a089677cf87E
                    local.get 1
                    return
                  end
                  local.get 0
                  i32.load offset=404
                  local.get 7
                  i32.add
                  local.tee 7
                  local.get 4
                  i32.le_u
                  br_if 1 (;@6;)
                  local.get 5
                  local.get 4
                  local.get 6
                  i32.const 1
                  i32.and
                  i32.or
                  i32.const 2
                  i32.or
                  i32.store
                  local.get 8
                  local.get 4
                  i32.add
                  local.tee 2
                  local.get 7
                  local.get 4
                  i32.sub
                  local.tee 3
                  i32.const 1
                  i32.or
                  i32.store offset=4
                  local.get 0
                  i32.const 404
                  i32.add
                  local.get 3
                  i32.store
                  local.get 0
                  i32.const 412
                  i32.add
                  local.get 2
                  i32.store
                  local.get 1
                  return
                end
                local.get 0
                i32.load offset=400
                local.get 7
                i32.add
                local.tee 7
                local.get 4
                i32.ge_u
                br_if 2 (;@4;)
              end
              local.get 0
              local.get 2
              call $_ZN8dlmalloc8dlmalloc8Dlmalloc6malloc17h7ef7d8a98d4afe8dE
              local.tee 4
              i32.eqz
              br_if 0 (;@5;)
              local.get 4
              local.get 1
              local.get 2
              local.get 5
              i32.load
              local.tee 3
              i32.const -8
              i32.and
              i32.const 4
              i32.const 8
              local.get 3
              i32.const 3
              i32.and
              select
              i32.sub
              local.tee 3
              local.get 3
              local.get 2
              i32.gt_u
              select
              call $memcpy
              local.set 2
              local.get 0
              local.get 1
              call $_ZN8dlmalloc8dlmalloc8Dlmalloc4free17he25d64d2cc54b15fE
              local.get 2
              local.set 3
            end
            local.get 3
            return
          end
          block  ;; label = @4
            block  ;; label = @5
              local.get 7
              local.get 4
              i32.sub
              local.tee 2
              i32.const 16
              i32.ge_u
              br_if 0 (;@5;)
              local.get 5
              local.get 6
              i32.const 1
              i32.and
              local.get 7
              i32.or
              i32.const 2
              i32.or
              i32.store
              local.get 8
              local.get 7
              i32.add
              local.tee 2
              local.get 2
              i32.load offset=4
              i32.const 1
              i32.or
              i32.store offset=4
              i32.const 0
              local.set 2
              i32.const 0
              local.set 3
              br 1 (;@4;)
            end
            local.get 5
            local.get 4
            local.get 6
            i32.const 1
            i32.and
            i32.or
            i32.const 2
            i32.or
            i32.store
            local.get 8
            local.get 4
            i32.add
            local.tee 3
            local.get 2
            i32.const 1
            i32.or
            i32.store offset=4
            local.get 8
            local.get 7
            i32.add
            local.tee 4
            local.get 2
            i32.store
            local.get 4
            local.get 4
            i32.load offset=4
            i32.const -2
            i32.and
            i32.store offset=4
          end
          local.get 0
          i32.const 408
          i32.add
          local.get 3
          i32.store
          local.get 0
          i32.const 400
          i32.add
          local.get 2
          i32.store
          local.get 1
          return
        end
        local.get 0
        local.get 9
        call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
        br 1 (;@1;)
      end
      local.get 0
      local.get 0
      i32.load
      i32.const -2
      local.get 6
      i32.const 3
      i32.shr_u
      i32.rotl
      i32.and
      i32.store
    end
    block  ;; label = @1
      local.get 2
      i32.const 15
      i32.gt_u
      br_if 0 (;@1;)
      local.get 5
      local.get 7
      local.get 5
      i32.load
      i32.const 1
      i32.and
      i32.or
      i32.const 2
      i32.or
      i32.store
      local.get 8
      local.get 7
      i32.add
      local.tee 2
      local.get 2
      i32.load offset=4
      i32.const 1
      i32.or
      i32.store offset=4
      local.get 1
      return
    end
    local.get 5
    local.get 4
    local.get 5
    i32.load
    i32.const 1
    i32.and
    i32.or
    i32.const 2
    i32.or
    i32.store
    local.get 8
    local.get 4
    i32.add
    local.tee 3
    local.get 2
    i32.const 3
    i32.or
    i32.store offset=4
    local.get 8
    local.get 7
    i32.add
    local.tee 4
    local.get 4
    i32.load offset=4
    i32.const 1
    i32.or
    i32.store offset=4
    local.get 0
    local.get 3
    local.get 2
    call $_ZN8dlmalloc8dlmalloc8Dlmalloc13dispose_chunk17h08c63a089677cf87E
    local.get 1)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc13dispose_chunk17h08c63a089677cf87E (type 8) (param i32 i32 i32)
    (local i32 i32 i32 i32)
    local.get 1
    local.get 2
    i32.add
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 1
                    i32.load offset=4
                    local.tee 4
                    i32.const 1
                    i32.and
                    br_if 0 (;@8;)
                    local.get 4
                    i32.const 3
                    i32.and
                    i32.eqz
                    br_if 1 (;@7;)
                    local.get 1
                    i32.load
                    local.tee 4
                    local.get 2
                    i32.add
                    local.set 2
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 0
                          i32.load offset=408
                          local.get 1
                          local.get 4
                          i32.sub
                          local.tee 1
                          i32.eq
                          br_if 0 (;@11;)
                          local.get 4
                          i32.const 255
                          i32.gt_u
                          br_if 1 (;@10;)
                          local.get 1
                          i32.load offset=12
                          local.tee 5
                          local.get 1
                          i32.load offset=8
                          local.tee 6
                          i32.eq
                          br_if 2 (;@9;)
                          local.get 6
                          local.get 5
                          i32.store offset=12
                          local.get 5
                          local.get 6
                          i32.store offset=8
                          br 3 (;@8;)
                        end
                        local.get 3
                        i32.load offset=4
                        i32.const 3
                        i32.and
                        i32.const 3
                        i32.ne
                        br_if 2 (;@8;)
                        local.get 0
                        local.get 2
                        i32.store offset=400
                        local.get 3
                        i32.const 4
                        i32.add
                        local.tee 0
                        local.get 0
                        i32.load
                        i32.const -2
                        i32.and
                        i32.store
                        local.get 1
                        local.get 2
                        i32.const 1
                        i32.or
                        i32.store offset=4
                        local.get 3
                        local.get 2
                        i32.store
                        return
                      end
                      local.get 0
                      local.get 1
                      call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
                      br 1 (;@8;)
                    end
                    local.get 0
                    local.get 0
                    i32.load
                    i32.const -2
                    local.get 4
                    i32.const 3
                    i32.shr_u
                    i32.rotl
                    i32.and
                    i32.store
                  end
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 3
                      i32.load offset=4
                      local.tee 4
                      i32.const 2
                      i32.and
                      br_if 0 (;@9;)
                      local.get 0
                      i32.load offset=412
                      local.get 3
                      i32.eq
                      br_if 1 (;@8;)
                      local.get 0
                      i32.load offset=408
                      local.get 3
                      i32.eq
                      br_if 3 (;@6;)
                      local.get 4
                      i32.const -8
                      i32.and
                      local.tee 5
                      local.get 2
                      i32.add
                      local.set 2
                      local.get 5
                      i32.const 255
                      i32.gt_u
                      br_if 4 (;@5;)
                      local.get 3
                      i32.load offset=12
                      local.tee 5
                      local.get 3
                      i32.load offset=8
                      local.tee 3
                      i32.eq
                      br_if 6 (;@3;)
                      local.get 3
                      local.get 5
                      i32.store offset=12
                      local.get 5
                      local.get 3
                      i32.store offset=8
                      br 7 (;@2;)
                    end
                    local.get 3
                    i32.const 4
                    i32.add
                    local.get 4
                    i32.const -2
                    i32.and
                    i32.store
                    local.get 1
                    local.get 2
                    i32.const 1
                    i32.or
                    i32.store offset=4
                    local.get 1
                    local.get 2
                    i32.add
                    local.get 2
                    i32.store
                    br 7 (;@1;)
                  end
                  local.get 0
                  i32.const 412
                  i32.add
                  local.get 1
                  i32.store
                  local.get 0
                  local.get 0
                  i32.load offset=404
                  local.get 2
                  i32.add
                  local.tee 2
                  i32.store offset=404
                  local.get 1
                  local.get 2
                  i32.const 1
                  i32.or
                  i32.store offset=4
                  local.get 1
                  local.get 0
                  i32.load offset=408
                  i32.eq
                  br_if 3 (;@4;)
                end
                return
              end
              local.get 0
              i32.const 408
              i32.add
              local.get 1
              i32.store
              local.get 0
              local.get 0
              i32.load offset=400
              local.get 2
              i32.add
              local.tee 2
              i32.store offset=400
              local.get 1
              local.get 2
              i32.const 1
              i32.or
              i32.store offset=4
              local.get 1
              local.get 2
              i32.add
              local.get 2
              i32.store
              return
            end
            local.get 0
            local.get 3
            call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
            br 2 (;@2;)
          end
          local.get 0
          i32.const 0
          i32.store offset=400
          local.get 0
          i32.const 408
          i32.add
          i32.const 0
          i32.store
          return
        end
        local.get 0
        local.get 0
        i32.load
        i32.const -2
        local.get 4
        i32.const 3
        i32.shr_u
        i32.rotl
        i32.and
        i32.store
      end
      local.get 1
      local.get 2
      i32.const 1
      i32.or
      i32.store offset=4
      local.get 1
      local.get 2
      i32.add
      local.get 2
      i32.store
      local.get 1
      local.get 0
      i32.const 408
      i32.add
      i32.load
      i32.ne
      br_if 0 (;@1;)
      local.get 0
      local.get 2
      i32.store offset=400
      return
    end
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          i32.const 255
          i32.gt_u
          br_if 0 (;@3;)
          local.get 0
          local.get 2
          i32.const 3
          i32.shr_u
          local.tee 3
          i32.const 3
          i32.shl
          i32.add
          i32.const 8
          i32.add
          local.set 2
          local.get 0
          i32.load
          local.tee 4
          i32.const 1
          local.get 3
          i32.const 31
          i32.and
          i32.shl
          local.tee 3
          i32.and
          i32.eqz
          br_if 1 (;@2;)
          local.get 2
          i32.load offset=8
          local.set 0
          br 2 (;@1;)
        end
        local.get 0
        local.get 1
        local.get 2
        call $_ZN8dlmalloc8dlmalloc8Dlmalloc18insert_large_chunk17he00600b9254d4992E
        return
      end
      local.get 0
      local.get 4
      local.get 3
      i32.or
      i32.store
      local.get 2
      local.set 0
    end
    local.get 2
    local.get 1
    i32.store offset=8
    local.get 0
    local.get 1
    i32.store offset=12
    local.get 1
    local.get 2
    i32.store offset=12
    local.get 1
    local.get 0
    i32.store offset=8)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc4free17he25d64d2cc54b15fE (type 1) (param i32 i32)
    (local i32 i32 i32 i32 i32)
    local.get 1
    i32.const -8
    i32.add
    local.tee 2
    local.get 1
    i32.const -4
    i32.add
    i32.load
    local.tee 3
    i32.const -8
    i32.and
    local.tee 1
    i32.add
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        local.get 3
        i32.const 1
        i32.and
        br_if 0 (;@2;)
        local.get 3
        i32.const 3
        i32.and
        i32.eqz
        br_if 1 (;@1;)
        local.get 2
        i32.load
        local.tee 3
        local.get 1
        i32.add
        local.set 1
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.load offset=408
              local.get 2
              local.get 3
              i32.sub
              local.tee 2
              i32.eq
              br_if 0 (;@5;)
              local.get 3
              i32.const 255
              i32.gt_u
              br_if 1 (;@4;)
              local.get 2
              i32.load offset=12
              local.tee 5
              local.get 2
              i32.load offset=8
              local.tee 6
              i32.eq
              br_if 2 (;@3;)
              local.get 6
              local.get 5
              i32.store offset=12
              local.get 5
              local.get 6
              i32.store offset=8
              br 3 (;@2;)
            end
            local.get 4
            i32.load offset=4
            i32.const 3
            i32.and
            i32.const 3
            i32.ne
            br_if 2 (;@2;)
            local.get 0
            local.get 1
            i32.store offset=400
            local.get 4
            i32.const 4
            i32.add
            local.tee 0
            local.get 0
            i32.load
            i32.const -2
            i32.and
            i32.store
            local.get 2
            local.get 1
            i32.const 1
            i32.or
            i32.store offset=4
            local.get 2
            local.get 1
            i32.add
            local.get 1
            i32.store
            return
          end
          local.get 0
          local.get 2
          call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
          br 1 (;@2;)
        end
        local.get 0
        local.get 0
        i32.load
        i32.const -2
        local.get 3
        i32.const 3
        i32.shr_u
        i32.rotl
        i32.and
        i32.store
      end
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 4
                        i32.load offset=4
                        local.tee 3
                        i32.const 2
                        i32.and
                        br_if 0 (;@10;)
                        local.get 0
                        i32.load offset=412
                        local.get 4
                        i32.eq
                        br_if 1 (;@9;)
                        local.get 0
                        i32.load offset=408
                        local.get 4
                        i32.eq
                        br_if 2 (;@8;)
                        local.get 3
                        i32.const -8
                        i32.and
                        local.tee 5
                        local.get 1
                        i32.add
                        local.set 1
                        local.get 5
                        i32.const 255
                        i32.gt_u
                        br_if 3 (;@7;)
                        local.get 4
                        i32.load offset=12
                        local.tee 5
                        local.get 4
                        i32.load offset=8
                        local.tee 4
                        i32.eq
                        br_if 4 (;@6;)
                        local.get 4
                        local.get 5
                        i32.store offset=12
                        local.get 5
                        local.get 4
                        i32.store offset=8
                        br 5 (;@5;)
                      end
                      local.get 4
                      i32.const 4
                      i32.add
                      local.get 3
                      i32.const -2
                      i32.and
                      i32.store
                      local.get 2
                      local.get 1
                      i32.const 1
                      i32.or
                      i32.store offset=4
                      local.get 2
                      local.get 1
                      i32.add
                      local.get 1
                      i32.store
                      br 7 (;@2;)
                    end
                    local.get 0
                    i32.const 412
                    i32.add
                    local.get 2
                    i32.store
                    local.get 0
                    local.get 0
                    i32.load offset=404
                    local.get 1
                    i32.add
                    local.tee 1
                    i32.store offset=404
                    local.get 2
                    local.get 1
                    i32.const 1
                    i32.or
                    i32.store offset=4
                    block  ;; label = @9
                      local.get 2
                      local.get 0
                      i32.load offset=408
                      i32.ne
                      br_if 0 (;@9;)
                      local.get 0
                      i32.const 0
                      i32.store offset=400
                      local.get 0
                      i32.const 408
                      i32.add
                      i32.const 0
                      i32.store
                    end
                    local.get 0
                    i32.load offset=440
                    local.tee 3
                    local.get 1
                    i32.ge_u
                    br_if 7 (;@1;)
                    local.get 0
                    i32.const 412
                    i32.add
                    i32.load
                    local.tee 1
                    i32.eqz
                    br_if 7 (;@1;)
                    block  ;; label = @9
                      local.get 0
                      i32.const 404
                      i32.add
                      i32.load
                      local.tee 5
                      i32.const 41
                      i32.lt_u
                      br_if 0 (;@9;)
                      local.get 0
                      i32.const 424
                      i32.add
                      local.set 2
                      loop  ;; label = @10
                        block  ;; label = @11
                          local.get 2
                          i32.load
                          local.tee 4
                          local.get 1
                          i32.gt_u
                          br_if 0 (;@11;)
                          local.get 4
                          local.get 2
                          i32.load offset=4
                          i32.add
                          local.get 1
                          i32.gt_u
                          br_if 2 (;@9;)
                        end
                        local.get 2
                        i32.load offset=8
                        local.tee 2
                        br_if 0 (;@10;)
                      end
                    end
                    local.get 0
                    i32.const 432
                    i32.add
                    i32.load
                    local.tee 1
                    i32.eqz
                    br_if 4 (;@4;)
                    i32.const 0
                    local.set 2
                    loop  ;; label = @9
                      local.get 2
                      i32.const 1
                      i32.add
                      local.set 2
                      local.get 1
                      i32.load offset=8
                      local.tee 1
                      br_if 0 (;@9;)
                    end
                    local.get 2
                    i32.const 4095
                    local.get 2
                    i32.const 4095
                    i32.gt_u
                    select
                    local.set 2
                    br 5 (;@3;)
                  end
                  local.get 0
                  i32.const 408
                  i32.add
                  local.get 2
                  i32.store
                  local.get 0
                  local.get 0
                  i32.load offset=400
                  local.get 1
                  i32.add
                  local.tee 1
                  i32.store offset=400
                  local.get 2
                  local.get 1
                  i32.const 1
                  i32.or
                  i32.store offset=4
                  local.get 2
                  local.get 1
                  i32.add
                  local.get 1
                  i32.store
                  return
                end
                local.get 0
                local.get 4
                call $_ZN8dlmalloc8dlmalloc8Dlmalloc18unlink_large_chunk17h5160d944f49d66b9E
                br 1 (;@5;)
              end
              local.get 0
              local.get 0
              i32.load
              i32.const -2
              local.get 3
              i32.const 3
              i32.shr_u
              i32.rotl
              i32.and
              i32.store
            end
            local.get 2
            local.get 1
            i32.const 1
            i32.or
            i32.store offset=4
            local.get 2
            local.get 1
            i32.add
            local.get 1
            i32.store
            local.get 2
            local.get 0
            i32.const 408
            i32.add
            i32.load
            i32.ne
            br_if 2 (;@2;)
            local.get 0
            local.get 1
            i32.store offset=400
            return
          end
          i32.const 4095
          local.set 2
        end
        local.get 0
        local.get 2
        i32.store offset=448
        local.get 5
        local.get 3
        i32.le_u
        br_if 1 (;@1;)
        local.get 0
        i32.const 440
        i32.add
        i32.const -1
        i32.store
        return
      end
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 1
                i32.const 255
                i32.gt_u
                br_if 0 (;@6;)
                local.get 0
                local.get 1
                i32.const 3
                i32.shr_u
                local.tee 4
                i32.const 3
                i32.shl
                i32.add
                i32.const 8
                i32.add
                local.set 1
                local.get 0
                i32.load
                local.tee 3
                i32.const 1
                local.get 4
                i32.const 31
                i32.and
                i32.shl
                local.tee 4
                i32.and
                i32.eqz
                br_if 1 (;@5;)
                local.get 1
                i32.load offset=8
                local.set 0
                br 2 (;@4;)
              end
              local.get 0
              local.get 2
              local.get 1
              call $_ZN8dlmalloc8dlmalloc8Dlmalloc18insert_large_chunk17he00600b9254d4992E
              local.get 0
              local.get 0
              i32.load offset=448
              i32.const -1
              i32.add
              local.tee 2
              i32.store offset=448
              local.get 2
              br_if 4 (;@1;)
              local.get 0
              i32.const 432
              i32.add
              i32.load
              local.tee 1
              i32.eqz
              br_if 2 (;@3;)
              i32.const 0
              local.set 2
              loop  ;; label = @6
                local.get 2
                i32.const 1
                i32.add
                local.set 2
                local.get 1
                i32.load offset=8
                local.tee 1
                br_if 0 (;@6;)
              end
              local.get 2
              i32.const 4095
              local.get 2
              i32.const 4095
              i32.gt_u
              select
              local.set 2
              br 3 (;@2;)
            end
            local.get 0
            local.get 3
            local.get 4
            i32.or
            i32.store
            local.get 1
            local.set 0
          end
          local.get 1
          local.get 2
          i32.store offset=8
          local.get 0
          local.get 2
          i32.store offset=12
          local.get 2
          local.get 1
          i32.store offset=12
          local.get 2
          local.get 0
          i32.store offset=8
          return
        end
        i32.const 4095
        local.set 2
      end
      local.get 0
      i32.const 448
      i32.add
      local.get 2
      i32.store
    end)
  (func $_ZN8dlmalloc8dlmalloc8Dlmalloc8memalign17hfb05805b97584ef4E (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32)
    i32.const 0
    local.set 3
    block  ;; label = @1
      i32.const -65587
      local.get 1
      i32.const 16
      local.get 1
      i32.const 16
      i32.gt_u
      select
      local.tee 1
      i32.sub
      local.get 2
      i32.le_u
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.const 16
      local.get 2
      i32.const 11
      i32.add
      i32.const -8
      i32.and
      local.get 2
      i32.const 11
      i32.lt_u
      select
      local.tee 4
      i32.add
      i32.const 12
      i32.add
      call $_ZN8dlmalloc8dlmalloc8Dlmalloc6malloc17h7ef7d8a98d4afe8dE
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      i32.const -8
      i32.add
      local.set 3
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.const -1
            i32.add
            local.tee 5
            local.get 2
            i32.and
            i32.eqz
            br_if 0 (;@4;)
            local.get 2
            i32.const -4
            i32.add
            local.tee 6
            i32.load
            local.tee 7
            i32.const -8
            i32.and
            local.get 5
            local.get 2
            i32.add
            i32.const 0
            local.get 1
            i32.sub
            i32.and
            i32.const -8
            i32.add
            local.tee 2
            local.get 2
            local.get 1
            i32.add
            local.get 2
            local.get 3
            i32.sub
            i32.const 16
            i32.gt_u
            select
            local.tee 1
            local.get 3
            i32.sub
            local.tee 2
            i32.sub
            local.set 5
            local.get 7
            i32.const 3
            i32.and
            i32.eqz
            br_if 1 (;@3;)
            local.get 1
            local.get 5
            local.get 1
            i32.load offset=4
            i32.const 1
            i32.and
            i32.or
            i32.const 2
            i32.or
            i32.store offset=4
            local.get 1
            local.get 5
            i32.add
            local.tee 5
            local.get 5
            i32.load offset=4
            i32.const 1
            i32.or
            i32.store offset=4
            local.get 6
            local.get 2
            local.get 6
            i32.load
            i32.const 1
            i32.and
            i32.or
            i32.const 2
            i32.or
            i32.store
            local.get 1
            local.get 1
            i32.load offset=4
            i32.const 1
            i32.or
            i32.store offset=4
            local.get 0
            local.get 3
            local.get 2
            call $_ZN8dlmalloc8dlmalloc8Dlmalloc13dispose_chunk17h08c63a089677cf87E
            br 2 (;@2;)
          end
          local.get 3
          local.set 1
          br 1 (;@2;)
        end
        local.get 3
        i32.load
        local.set 3
        local.get 1
        local.get 5
        i32.store offset=4
        local.get 1
        local.get 3
        local.get 2
        i32.add
        i32.store
      end
      block  ;; label = @2
        local.get 1
        i32.load offset=4
        local.tee 2
        i32.const 3
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        i32.const -8
        i32.and
        local.tee 3
        local.get 4
        i32.const 16
        i32.add
        i32.le_u
        br_if 0 (;@2;)
        local.get 1
        i32.const 4
        i32.add
        local.get 4
        local.get 2
        i32.const 1
        i32.and
        i32.or
        i32.const 2
        i32.or
        i32.store
        local.get 1
        local.get 4
        i32.add
        local.tee 2
        local.get 3
        local.get 4
        i32.sub
        local.tee 4
        i32.const 3
        i32.or
        i32.store offset=4
        local.get 1
        local.get 3
        i32.add
        local.tee 3
        local.get 3
        i32.load offset=4
        i32.const 1
        i32.or
        i32.store offset=4
        local.get 0
        local.get 2
        local.get 4
        call $_ZN8dlmalloc8dlmalloc8Dlmalloc13dispose_chunk17h08c63a089677cf87E
      end
      local.get 1
      i32.const 8
      i32.add
      local.set 3
    end
    local.get 3)
  (func $_ZN5alloc5alloc18handle_alloc_error17hf7e8102cebbd3235E (type 1) (param i32 i32)
    local.get 0
    local.get 1
    call $rust_oom
    unreachable)
  (func $_ZN5alloc7raw_vec17capacity_overflow17hd685e916963b651dE (type 5)
    i32.const 1050156
    call $_ZN4core9panicking5panic17h62fdcfa056e70982E
    unreachable)
  (func $_ZN5alloc6string104_$LT$impl$u20$core..convert..From$LT$alloc..string..String$GT$$u20$for$u20$alloc..vec..Vec$LT$u8$GT$$GT$4from17he2b8ded24866ca30E (type 1) (param i32 i32)
    local.get 0
    local.get 1
    i64.load align=4
    i64.store align=4
    local.get 0
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i32.load
    i32.store)
  (func $_ZN4core3ptr18real_drop_in_place17h0013f08692be52e6E (type 0) (param i32))
  (func $_ZN4core3ptr18real_drop_in_place17h23040579a46f15a6E (type 0) (param i32))
  (func $_ZN4core3ptr18real_drop_in_place17h5a9a7fbf25605767E (type 0) (param i32))
  (func $_ZN4core5slice20slice_index_len_fail17hb115deb2b20f49d8E (type 1) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.store offset=4
    local.get 2
    local.get 0
    i32.store
    local.get 2
    i32.const 44
    i32.add
    i32.const 30
    i32.store
    local.get 2
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 2
    i32.const 30
    i32.store offset=36
    local.get 2
    i64.const 2
    i64.store offset=12 align=4
    local.get 2
    i32.const 1050460
    i32.store offset=8
    local.get 2
    local.get 2
    i32.const 4
    i32.add
    i32.store offset=40
    local.get 2
    local.get 2
    i32.store offset=32
    local.get 2
    local.get 2
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 2
    i32.const 8
    i32.add
    i32.const 1050476
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E (type 8) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    local.get 2
    i32.store offset=4
    local.get 3
    local.get 1
    i32.store
    local.get 3
    i32.const 44
    i32.add
    i32.const 30
    i32.store
    local.get 3
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 30
    i32.store offset=36
    local.get 3
    i64.const 2
    i64.store offset=12 align=4
    local.get 3
    i32.const 1050292
    i32.store offset=8
    local.get 3
    local.get 3
    i32.store offset=40
    local.get 3
    local.get 3
    i32.const 4
    i32.add
    i32.store offset=32
    local.get 3
    local.get 3
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 3
    i32.const 8
    i32.add
    local.get 0
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN4core9panicking5panic17h62fdcfa056e70982E (type 0) (param i32)
    (local i32 i64 i64 i64)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    i64.load offset=8 align=4
    local.set 2
    local.get 0
    i64.load offset=16 align=4
    local.set 3
    local.get 0
    i64.load align=4
    local.set 4
    local.get 1
    i32.const 20
    i32.add
    i32.const 0
    i32.store
    local.get 1
    local.get 4
    i64.store offset=24
    local.get 1
    i32.const 1050184
    i32.store offset=16
    local.get 1
    i64.const 1
    i64.store offset=4 align=4
    local.get 1
    local.get 1
    i32.const 24
    i32.add
    i32.store
    local.get 1
    local.get 3
    i64.store offset=40
    local.get 1
    local.get 2
    i64.store offset=32
    local.get 1
    local.get 1
    i32.const 32
    i32.add
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E (type 1) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.store offset=4
    local.get 2
    local.get 0
    i32.store
    local.get 2
    i32.const 44
    i32.add
    i32.const 30
    i32.store
    local.get 2
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 2
    i32.const 30
    i32.store offset=36
    local.get 2
    i64.const 2
    i64.store offset=12 align=4
    local.get 2
    i32.const 1050528
    i32.store offset=8
    local.get 2
    local.get 2
    i32.const 4
    i32.add
    i32.store offset=40
    local.get 2
    local.get 2
    i32.store offset=32
    local.get 2
    local.get 2
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 2
    i32.const 8
    i32.add
    i32.const 1050544
    call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
    unreachable)
  (func $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE (type 1) (param i32 i32)
    (local i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 1
    i64.load align=4
    local.set 3
    local.get 2
    i32.const 20
    i32.add
    local.get 1
    i64.load offset=8 align=4
    i64.store align=4
    local.get 2
    local.get 3
    i64.store offset=12 align=4
    local.get 2
    local.get 0
    i32.store offset=8
    local.get 2
    i32.const 1050224
    i32.store offset=4
    local.get 2
    i32.const 1050184
    i32.store
    local.get 2
    call $rust_begin_unwind
    unreachable)
  (func $_ZN4core3fmt9Formatter3pad17hacaebd5abd28adf1E (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    local.get 0
    i32.load offset=16
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 0
                      i32.load offset=8
                      local.tee 4
                      i32.const 1
                      i32.ne
                      br_if 0 (;@9;)
                      local.get 3
                      br_if 1 (;@8;)
                      br 6 (;@3;)
                    end
                    local.get 3
                    i32.eqz
                    br_if 1 (;@7;)
                  end
                  local.get 2
                  i32.eqz
                  br_if 1 (;@6;)
                  local.get 1
                  local.get 2
                  i32.add
                  local.set 5
                  local.get 0
                  i32.const 20
                  i32.add
                  i32.load
                  i32.const -1
                  i32.xor
                  local.set 6
                  i32.const 0
                  local.set 7
                  local.get 1
                  local.set 3
                  local.get 1
                  local.set 8
                  loop  ;; label = @8
                    local.get 3
                    i32.const 1
                    i32.add
                    local.set 9
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              local.get 3
                              i32.load8_s
                              local.tee 10
                              i32.const 0
                              i32.lt_s
                              br_if 0 (;@13;)
                              local.get 10
                              i32.const 255
                              i32.and
                              local.set 10
                              br 1 (;@12;)
                            end
                            block  ;; label = @13
                              block  ;; label = @14
                                local.get 9
                                local.get 5
                                i32.eq
                                br_if 0 (;@14;)
                                local.get 9
                                i32.load8_u
                                i32.const 63
                                i32.and
                                local.set 11
                                local.get 3
                                i32.const 2
                                i32.add
                                local.tee 3
                                local.set 9
                                br 1 (;@13;)
                              end
                              i32.const 0
                              local.set 11
                              local.get 5
                              local.set 3
                            end
                            local.get 10
                            i32.const 31
                            i32.and
                            local.set 12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  local.get 10
                                  i32.const 255
                                  i32.and
                                  local.tee 10
                                  i32.const 224
                                  i32.lt_u
                                  br_if 0 (;@15;)
                                  local.get 3
                                  local.get 5
                                  i32.eq
                                  br_if 1 (;@14;)
                                  local.get 3
                                  i32.load8_u
                                  i32.const 63
                                  i32.and
                                  local.set 13
                                  local.get 3
                                  i32.const 1
                                  i32.add
                                  local.tee 9
                                  local.set 14
                                  br 2 (;@13;)
                                end
                                local.get 11
                                local.get 12
                                i32.const 6
                                i32.shl
                                i32.or
                                local.set 10
                                br 2 (;@12;)
                              end
                              i32.const 0
                              local.set 13
                              local.get 5
                              local.set 14
                            end
                            local.get 13
                            local.get 11
                            i32.const 6
                            i32.shl
                            i32.or
                            local.set 11
                            block  ;; label = @13
                              local.get 10
                              i32.const 240
                              i32.lt_u
                              br_if 0 (;@13;)
                              local.get 14
                              local.get 5
                              i32.eq
                              br_if 2 (;@11;)
                              local.get 14
                              i32.const 1
                              i32.add
                              local.set 3
                              local.get 14
                              i32.load8_u
                              i32.const 63
                              i32.and
                              local.set 10
                              br 3 (;@10;)
                            end
                            local.get 11
                            local.get 12
                            i32.const 12
                            i32.shl
                            i32.or
                            local.set 10
                          end
                          local.get 9
                          local.set 3
                          local.get 6
                          i32.const 1
                          i32.add
                          local.tee 6
                          br_if 2 (;@9;)
                          br 6 (;@5;)
                        end
                        i32.const 0
                        local.set 10
                        local.get 9
                        local.set 3
                      end
                      local.get 11
                      i32.const 6
                      i32.shl
                      local.get 12
                      i32.const 18
                      i32.shl
                      i32.const 1835008
                      i32.and
                      i32.or
                      local.get 10
                      i32.or
                      local.tee 10
                      i32.const 1114112
                      i32.eq
                      br_if 5 (;@4;)
                      local.get 6
                      i32.const 1
                      i32.add
                      local.tee 6
                      i32.eqz
                      br_if 4 (;@5;)
                    end
                    local.get 7
                    local.get 8
                    i32.sub
                    local.get 3
                    i32.add
                    local.set 7
                    local.get 3
                    local.set 8
                    local.get 5
                    local.get 3
                    i32.ne
                    br_if 0 (;@8;)
                    br 4 (;@4;)
                  end
                end
                local.get 0
                i32.load offset=24
                local.get 1
                local.get 2
                local.get 0
                i32.const 28
                i32.add
                i32.load
                i32.load offset=12
                call_indirect (type 3)
                local.set 3
                br 5 (;@1;)
              end
              i32.const 0
              local.set 2
              local.get 4
              br_if 2 (;@3;)
              br 3 (;@2;)
            end
            local.get 10
            i32.const 1114112
            i32.eq
            br_if 0 (;@4;)
            block  ;; label = @5
              block  ;; label = @6
                local.get 7
                i32.eqz
                br_if 0 (;@6;)
                local.get 7
                local.get 2
                i32.eq
                br_if 0 (;@6;)
                i32.const 0
                local.set 3
                local.get 7
                local.get 2
                i32.ge_u
                br_if 1 (;@5;)
                local.get 1
                local.get 7
                i32.add
                i32.load8_s
                i32.const -64
                i32.lt_s
                br_if 1 (;@5;)
              end
              local.get 1
              local.set 3
            end
            local.get 7
            local.get 2
            local.get 3
            select
            local.set 2
            local.get 3
            local.get 1
            local.get 3
            select
            local.set 1
          end
          local.get 4
          i32.eqz
          br_if 1 (;@2;)
        end
        i32.const 0
        local.set 9
        block  ;; label = @3
          local.get 2
          i32.eqz
          br_if 0 (;@3;)
          local.get 2
          local.set 10
          local.get 1
          local.set 3
          loop  ;; label = @4
            local.get 9
            local.get 3
            i32.load8_u
            i32.const 192
            i32.and
            i32.const 128
            i32.eq
            i32.add
            local.set 9
            local.get 3
            i32.const 1
            i32.add
            local.set 3
            local.get 10
            i32.const -1
            i32.add
            local.tee 10
            br_if 0 (;@4;)
          end
        end
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 2
                local.get 9
                i32.sub
                local.get 0
                i32.const 12
                i32.add
                i32.load
                local.tee 6
                i32.ge_u
                br_if 0 (;@6;)
                i32.const 0
                local.set 9
                block  ;; label = @7
                  local.get 2
                  i32.eqz
                  br_if 0 (;@7;)
                  i32.const 0
                  local.set 9
                  local.get 2
                  local.set 10
                  local.get 1
                  local.set 3
                  loop  ;; label = @8
                    local.get 9
                    local.get 3
                    i32.load8_u
                    i32.const 192
                    i32.and
                    i32.const 128
                    i32.eq
                    i32.add
                    local.set 9
                    local.get 3
                    i32.const 1
                    i32.add
                    local.set 3
                    local.get 10
                    i32.const -1
                    i32.add
                    local.tee 10
                    br_if 0 (;@8;)
                  end
                end
                local.get 9
                local.get 2
                i32.sub
                local.get 6
                i32.add
                local.set 9
                i32.const 0
                local.get 0
                i32.load8_u offset=48
                local.tee 3
                local.get 3
                i32.const 3
                i32.eq
                select
                local.tee 3
                i32.const 3
                i32.and
                i32.eqz
                br_if 1 (;@5;)
                local.get 3
                i32.const 2
                i32.eq
                br_if 2 (;@4;)
                i32.const 0
                local.set 8
                br 3 (;@3;)
              end
              local.get 0
              i32.load offset=24
              local.get 1
              local.get 2
              local.get 0
              i32.const 28
              i32.add
              i32.load
              i32.load offset=12
              call_indirect (type 3)
              return
            end
            local.get 9
            local.set 8
            i32.const 0
            local.set 9
            br 1 (;@3;)
          end
          local.get 9
          i32.const 1
          i32.add
          i32.const 1
          i32.shr_u
          local.set 8
          local.get 9
          i32.const 1
          i32.shr_u
          local.set 9
        end
        i32.const -1
        local.set 3
        local.get 0
        i32.const 4
        i32.add
        local.set 10
        local.get 0
        i32.const 24
        i32.add
        local.set 6
        local.get 0
        i32.const 28
        i32.add
        local.set 7
        block  ;; label = @3
          loop  ;; label = @4
            local.get 3
            i32.const 1
            i32.add
            local.tee 3
            local.get 9
            i32.ge_u
            br_if 1 (;@3;)
            local.get 6
            i32.load
            local.get 10
            i32.load
            local.get 7
            i32.load
            i32.load offset=16
            call_indirect (type 4)
            i32.eqz
            br_if 0 (;@4;)
          end
          i32.const 1
          return
        end
        local.get 0
        i32.const 4
        i32.add
        i32.load
        local.set 9
        i32.const 1
        local.set 3
        local.get 0
        i32.const 24
        i32.add
        local.tee 10
        i32.load
        local.get 1
        local.get 2
        local.get 0
        i32.const 28
        i32.add
        local.tee 6
        i32.load
        i32.load offset=12
        call_indirect (type 3)
        br_if 1 (;@1;)
        local.get 10
        i32.load
        local.set 10
        i32.const -1
        local.set 3
        local.get 6
        i32.load
        i32.const 16
        i32.add
        local.set 6
        block  ;; label = @3
          loop  ;; label = @4
            local.get 3
            i32.const 1
            i32.add
            local.tee 3
            local.get 8
            i32.ge_u
            br_if 1 (;@3;)
            local.get 10
            local.get 9
            local.get 6
            i32.load
            call_indirect (type 4)
            i32.eqz
            br_if 0 (;@4;)
          end
          i32.const 1
          return
        end
        i32.const 0
        return
      end
      local.get 0
      i32.load offset=24
      local.get 1
      local.get 2
      local.get 0
      i32.const 28
      i32.add
      i32.load
      i32.load offset=12
      call_indirect (type 3)
      return
    end
    local.get 3)
  (func $_ZN4core3str16slice_error_fail17h3704ce74b976be71E (type 9) (param i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 112
    i32.sub
    local.tee 4
    global.set 0
    local.get 4
    local.get 3
    i32.store offset=12
    local.get 4
    local.get 2
    i32.store offset=8
    i32.const 1
    local.set 5
    local.get 1
    local.set 6
    block  ;; label = @1
      local.get 1
      i32.const 257
      i32.lt_u
      br_if 0 (;@1;)
      i32.const 0
      local.get 1
      i32.sub
      local.set 7
      i32.const 256
      local.set 8
      block  ;; label = @2
        loop  ;; label = @3
          block  ;; label = @4
            local.get 8
            local.get 1
            i32.ge_u
            br_if 0 (;@4;)
            local.get 0
            local.get 8
            i32.add
            i32.load8_s
            i32.const -65
            i32.gt_s
            br_if 2 (;@2;)
          end
          local.get 8
          i32.const -1
          i32.add
          local.set 6
          i32.const 0
          local.set 5
          local.get 8
          i32.const 1
          i32.eq
          br_if 2 (;@1;)
          local.get 7
          local.get 8
          i32.add
          local.set 9
          local.get 6
          local.set 8
          local.get 9
          i32.const 1
          i32.ne
          br_if 0 (;@3;)
          br 2 (;@1;)
        end
      end
      i32.const 0
      local.set 5
      local.get 8
      local.set 6
    end
    local.get 4
    local.get 6
    i32.store offset=20
    local.get 4
    local.get 0
    i32.store offset=16
    local.get 4
    i32.const 0
    i32.const 5
    local.get 5
    select
    i32.store offset=28
    local.get 4
    i32.const 1050184
    i32.const 1050582
    local.get 5
    select
    i32.store offset=24
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          local.get 1
          i32.gt_u
          local.tee 8
          br_if 0 (;@3;)
          local.get 3
          local.get 1
          i32.gt_u
          br_if 0 (;@3;)
          local.get 2
          local.get 3
          i32.gt_u
          br_if 1 (;@2;)
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.eqz
              br_if 0 (;@5;)
              local.get 1
              local.get 2
              i32.eq
              br_if 0 (;@5;)
              local.get 1
              local.get 2
              i32.le_u
              br_if 1 (;@4;)
              local.get 0
              local.get 2
              i32.add
              i32.load8_s
              i32.const -64
              i32.lt_s
              br_if 1 (;@4;)
            end
            local.get 3
            local.set 2
          end
          local.get 4
          local.get 2
          i32.store offset=32
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.eqz
              br_if 0 (;@5;)
              local.get 2
              local.get 1
              i32.eq
              br_if 0 (;@5;)
              local.get 1
              i32.const 1
              i32.add
              local.set 9
              loop  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  local.get 1
                  i32.ge_u
                  br_if 0 (;@7;)
                  local.get 0
                  local.get 2
                  i32.add
                  i32.load8_s
                  i32.const -64
                  i32.ge_s
                  br_if 2 (;@5;)
                end
                local.get 2
                i32.const -1
                i32.add
                local.set 8
                local.get 2
                i32.const 1
                i32.eq
                br_if 2 (;@4;)
                local.get 9
                local.get 2
                i32.eq
                local.set 6
                local.get 8
                local.set 2
                local.get 6
                i32.eqz
                br_if 0 (;@6;)
                br 2 (;@4;)
              end
            end
            local.get 2
            local.set 8
          end
          local.get 8
          local.get 1
          i32.eq
          br_if 2 (;@1;)
          i32.const 1
          local.set 6
          i32.const 0
          local.set 5
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              local.get 8
              i32.add
              local.tee 9
              i32.load8_s
              local.tee 2
              i32.const 0
              i32.lt_s
              br_if 0 (;@5;)
              local.get 4
              local.get 2
              i32.const 255
              i32.and
              i32.store offset=36
              local.get 4
              i32.const 40
              i32.add
              local.set 2
              br 1 (;@4;)
            end
            local.get 0
            local.get 1
            i32.add
            local.tee 6
            local.set 1
            block  ;; label = @5
              local.get 9
              i32.const 1
              i32.add
              local.get 6
              i32.eq
              br_if 0 (;@5;)
              local.get 9
              i32.const 2
              i32.add
              local.set 1
              local.get 9
              i32.const 1
              i32.add
              i32.load8_u
              i32.const 63
              i32.and
              local.set 5
            end
            local.get 2
            i32.const 31
            i32.and
            local.set 9
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  i32.const 255
                  i32.and
                  i32.const 224
                  i32.lt_u
                  br_if 0 (;@7;)
                  i32.const 0
                  local.set 0
                  local.get 6
                  local.set 7
                  block  ;; label = @8
                    local.get 1
                    local.get 6
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 1
                    i32.const 1
                    i32.add
                    local.set 7
                    local.get 1
                    i32.load8_u
                    i32.const 63
                    i32.and
                    local.set 0
                  end
                  local.get 0
                  local.get 5
                  i32.const 6
                  i32.shl
                  i32.or
                  local.set 1
                  local.get 2
                  i32.const 255
                  i32.and
                  i32.const 240
                  i32.lt_u
                  br_if 1 (;@6;)
                  i32.const 0
                  local.set 2
                  block  ;; label = @8
                    local.get 7
                    local.get 6
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 7
                    i32.load8_u
                    i32.const 63
                    i32.and
                    local.set 2
                  end
                  local.get 1
                  i32.const 6
                  i32.shl
                  local.get 9
                  i32.const 18
                  i32.shl
                  i32.const 1835008
                  i32.and
                  i32.or
                  local.get 2
                  i32.or
                  local.tee 1
                  i32.const 1114112
                  i32.eq
                  br_if 6 (;@1;)
                  br 2 (;@5;)
                end
                local.get 5
                local.get 9
                i32.const 6
                i32.shl
                i32.or
                local.set 1
                br 1 (;@5;)
              end
              local.get 1
              local.get 9
              i32.const 12
              i32.shl
              i32.or
              local.set 1
            end
            local.get 4
            local.get 1
            i32.store offset=36
            i32.const 1
            local.set 6
            local.get 4
            i32.const 40
            i32.add
            local.set 2
            local.get 1
            i32.const 128
            i32.lt_u
            br_if 0 (;@4;)
            i32.const 2
            local.set 6
            local.get 1
            i32.const 2048
            i32.lt_u
            br_if 0 (;@4;)
            i32.const 3
            i32.const 4
            local.get 1
            i32.const 65536
            i32.lt_u
            select
            local.set 6
          end
          local.get 4
          local.get 8
          i32.store offset=40
          local.get 4
          local.get 6
          local.get 8
          i32.add
          i32.store offset=44
          local.get 4
          i32.const 108
          i32.add
          i32.const 31
          i32.store
          local.get 4
          i32.const 100
          i32.add
          i32.const 31
          i32.store
          local.get 4
          i32.const 72
          i32.add
          i32.const 20
          i32.add
          i32.const 32
          i32.store
          local.get 4
          i32.const 84
          i32.add
          i32.const 33
          i32.store
          local.get 4
          i32.const 48
          i32.add
          i32.const 20
          i32.add
          i32.const 5
          i32.store
          local.get 4
          local.get 2
          i32.store offset=88
          local.get 4
          i32.const 30
          i32.store offset=76
          local.get 4
          i64.const 5
          i64.store offset=52 align=4
          local.get 4
          i32.const 1050796
          i32.store offset=48
          local.get 4
          local.get 4
          i32.const 24
          i32.add
          i32.store offset=104
          local.get 4
          local.get 4
          i32.const 16
          i32.add
          i32.store offset=96
          local.get 4
          local.get 4
          i32.const 36
          i32.add
          i32.store offset=80
          local.get 4
          local.get 4
          i32.const 32
          i32.add
          i32.store offset=72
          local.get 4
          local.get 4
          i32.const 72
          i32.add
          i32.store offset=64
          local.get 4
          i32.const 48
          i32.add
          i32.const 1050836
          call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
          unreachable
        end
        local.get 4
        local.get 2
        local.get 3
        local.get 8
        select
        i32.store offset=40
        local.get 4
        i32.const 72
        i32.add
        i32.const 20
        i32.add
        i32.const 31
        i32.store
        local.get 4
        i32.const 84
        i32.add
        i32.const 31
        i32.store
        local.get 4
        i32.const 48
        i32.add
        i32.const 20
        i32.add
        i32.const 3
        i32.store
        local.get 4
        i32.const 30
        i32.store offset=76
        local.get 4
        i64.const 3
        i64.store offset=52 align=4
        local.get 4
        i32.const 1050620
        i32.store offset=48
        local.get 4
        local.get 4
        i32.const 24
        i32.add
        i32.store offset=88
        local.get 4
        local.get 4
        i32.const 16
        i32.add
        i32.store offset=80
        local.get 4
        local.get 4
        i32.const 40
        i32.add
        i32.store offset=72
        local.get 4
        local.get 4
        i32.const 72
        i32.add
        i32.store offset=64
        local.get 4
        i32.const 48
        i32.add
        i32.const 1050644
        call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
        unreachable
      end
      local.get 4
      i32.const 100
      i32.add
      i32.const 31
      i32.store
      local.get 4
      i32.const 72
      i32.add
      i32.const 20
      i32.add
      i32.const 31
      i32.store
      local.get 4
      i32.const 84
      i32.add
      i32.const 30
      i32.store
      local.get 4
      i32.const 48
      i32.add
      i32.const 20
      i32.add
      i32.const 4
      i32.store
      local.get 4
      i32.const 30
      i32.store offset=76
      local.get 4
      i64.const 4
      i64.store offset=52 align=4
      local.get 4
      i32.const 1050696
      i32.store offset=48
      local.get 4
      local.get 4
      i32.const 24
      i32.add
      i32.store offset=96
      local.get 4
      local.get 4
      i32.const 16
      i32.add
      i32.store offset=88
      local.get 4
      local.get 4
      i32.const 12
      i32.add
      i32.store offset=80
      local.get 4
      local.get 4
      i32.const 8
      i32.add
      i32.store offset=72
      local.get 4
      local.get 4
      i32.const 72
      i32.add
      i32.store offset=64
      local.get 4
      i32.const 48
      i32.add
      i32.const 1050728
      call $_ZN4core9panicking9panic_fmt17hc4f83bfed80aeabdE
      unreachable
    end
    i32.const 1050372
    call $_ZN4core9panicking5panic17h62fdcfa056e70982E
    unreachable)
  (func $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$u32$GT$3fmt17h4f3293445fab7cb7E (type 4) (param i32 i32) (result i32)
    local.get 0
    i64.load32_u
    i32.const 1
    local.get 1
    call $_ZN4core3fmt3num3imp7fmt_u6417h4648b2300e699242E)
  (func $_ZN4core3fmt5write17h8cfd01c67a4a46c9E (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i32.const 36
    i32.add
    local.get 1
    i32.store
    local.get 3
    i32.const 52
    i32.add
    local.get 2
    i32.const 20
    i32.add
    i32.load
    local.tee 4
    i32.store
    local.get 3
    i32.const 3
    i32.store8 offset=56
    local.get 3
    i32.const 44
    i32.add
    local.get 2
    i32.load offset=16
    local.tee 5
    local.get 4
    i32.const 3
    i32.shl
    i32.add
    i32.store
    local.get 3
    i64.const 137438953472
    i64.store offset=8
    local.get 3
    local.get 0
    i32.store offset=32
    i32.const 0
    local.set 6
    local.get 3
    i32.const 0
    i32.store offset=24
    local.get 3
    i32.const 0
    i32.store offset=16
    local.get 3
    local.get 5
    i32.store offset=48
    local.get 3
    local.get 5
    i32.store offset=40
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.load offset=8
              local.tee 7
              i32.eqz
              br_if 0 (;@5;)
              local.get 2
              i32.load
              local.set 8
              local.get 2
              i32.load offset=4
              local.tee 9
              local.get 2
              i32.const 12
              i32.add
              i32.load
              local.tee 5
              local.get 5
              local.get 9
              i32.gt_u
              select
              local.tee 10
              i32.eqz
              br_if 1 (;@4;)
              local.get 0
              local.get 8
              i32.load
              local.get 8
              i32.load offset=4
              local.get 1
              i32.load offset=12
              call_indirect (type 3)
              br_if 2 (;@3;)
              local.get 8
              i32.const 12
              i32.add
              local.set 5
              local.get 3
              i32.const 56
              i32.add
              local.set 1
              local.get 3
              i32.const 52
              i32.add
              local.set 11
              local.get 3
              i32.const 48
              i32.add
              local.set 12
              i32.const 1
              local.set 6
              block  ;; label = @6
                loop  ;; label = @7
                  local.get 1
                  local.get 7
                  i32.const 32
                  i32.add
                  i32.load8_u
                  i32.store8
                  local.get 3
                  local.get 7
                  i32.const 8
                  i32.add
                  i32.load
                  i32.store offset=12
                  local.get 3
                  local.get 7
                  i32.const 12
                  i32.add
                  i32.load
                  i32.store offset=8
                  i32.const 0
                  local.set 2
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            local.get 7
                            i32.const 24
                            i32.add
                            i32.load
                            local.tee 0
                            i32.const 1
                            i32.eq
                            br_if 0 (;@12;)
                            block  ;; label = @13
                              local.get 0
                              i32.const 2
                              i32.eq
                              br_if 0 (;@13;)
                              local.get 0
                              i32.const 3
                              i32.eq
                              br_if 5 (;@8;)
                              local.get 7
                              i32.const 28
                              i32.add
                              i32.load
                              local.set 4
                              br 2 (;@11;)
                            end
                            local.get 3
                            i32.const 8
                            i32.add
                            i32.const 32
                            i32.add
                            local.tee 4
                            i32.load
                            local.tee 0
                            local.get 3
                            i32.const 8
                            i32.add
                            i32.const 36
                            i32.add
                            i32.load
                            i32.eq
                            br_if 2 (;@10;)
                            local.get 4
                            local.get 0
                            i32.const 8
                            i32.add
                            i32.store
                            local.get 0
                            i32.load offset=4
                            i32.const 34
                            i32.ne
                            br_if 4 (;@8;)
                            local.get 0
                            i32.load
                            i32.load
                            local.set 4
                            br 1 (;@11;)
                          end
                          local.get 7
                          i32.const 28
                          i32.add
                          i32.load
                          local.tee 0
                          local.get 11
                          i32.load
                          local.tee 4
                          i32.ge_u
                          br_if 2 (;@9;)
                          local.get 12
                          i32.load
                          local.get 0
                          i32.const 3
                          i32.shl
                          i32.add
                          local.tee 0
                          i32.load offset=4
                          i32.const 34
                          i32.ne
                          br_if 3 (;@8;)
                          local.get 0
                          i32.load
                          i32.load
                          local.set 4
                        end
                        i32.const 1
                        local.set 2
                        br 2 (;@8;)
                      end
                      br 1 (;@8;)
                    end
                    i32.const 1051160
                    local.get 0
                    local.get 4
                    call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
                    unreachable
                  end
                  local.get 3
                  i32.const 8
                  i32.add
                  i32.const 12
                  i32.add
                  local.get 4
                  i32.store
                  local.get 3
                  i32.const 8
                  i32.add
                  i32.const 8
                  i32.add
                  local.get 2
                  i32.store
                  i32.const 0
                  local.set 2
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            local.get 7
                            i32.const 16
                            i32.add
                            i32.load
                            local.tee 0
                            i32.const 1
                            i32.eq
                            br_if 0 (;@12;)
                            block  ;; label = @13
                              local.get 0
                              i32.const 2
                              i32.eq
                              br_if 0 (;@13;)
                              local.get 0
                              i32.const 3
                              i32.eq
                              br_if 5 (;@8;)
                              local.get 7
                              i32.const 20
                              i32.add
                              i32.load
                              local.set 4
                              br 2 (;@11;)
                            end
                            local.get 3
                            i32.const 8
                            i32.add
                            i32.const 32
                            i32.add
                            local.tee 4
                            i32.load
                            local.tee 0
                            local.get 3
                            i32.const 8
                            i32.add
                            i32.const 36
                            i32.add
                            i32.load
                            i32.eq
                            br_if 2 (;@10;)
                            local.get 4
                            local.get 0
                            i32.const 8
                            i32.add
                            i32.store
                            local.get 0
                            i32.load offset=4
                            i32.const 34
                            i32.ne
                            br_if 4 (;@8;)
                            local.get 0
                            i32.load
                            i32.load
                            local.set 4
                            br 1 (;@11;)
                          end
                          local.get 7
                          i32.const 20
                          i32.add
                          i32.load
                          local.tee 0
                          local.get 11
                          i32.load
                          local.tee 4
                          i32.ge_u
                          br_if 2 (;@9;)
                          local.get 12
                          i32.load
                          local.get 0
                          i32.const 3
                          i32.shl
                          i32.add
                          local.tee 0
                          i32.load offset=4
                          i32.const 34
                          i32.ne
                          br_if 3 (;@8;)
                          local.get 0
                          i32.load
                          i32.load
                          local.set 4
                        end
                        i32.const 1
                        local.set 2
                        br 2 (;@8;)
                      end
                      br 1 (;@8;)
                    end
                    i32.const 1051160
                    local.get 0
                    local.get 4
                    call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
                    unreachable
                  end
                  local.get 3
                  i32.const 8
                  i32.add
                  i32.const 20
                  i32.add
                  local.get 4
                  i32.store
                  local.get 3
                  i32.const 8
                  i32.add
                  i32.const 16
                  i32.add
                  local.get 2
                  i32.store
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 7
                        i32.load
                        i32.const 1
                        i32.ne
                        br_if 0 (;@10;)
                        local.get 7
                        i32.const 4
                        i32.add
                        i32.load
                        local.tee 2
                        local.get 11
                        i32.load
                        local.tee 4
                        i32.ge_u
                        br_if 2 (;@8;)
                        local.get 12
                        i32.load
                        local.get 2
                        i32.const 3
                        i32.shl
                        i32.add
                        local.set 2
                        br 1 (;@9;)
                      end
                      local.get 3
                      i32.const 8
                      i32.add
                      i32.const 32
                      i32.add
                      local.tee 4
                      i32.load
                      local.tee 2
                      local.get 3
                      i32.const 8
                      i32.add
                      i32.const 36
                      i32.add
                      i32.load
                      i32.eq
                      br_if 3 (;@6;)
                      local.get 4
                      local.get 2
                      i32.const 8
                      i32.add
                      i32.store
                    end
                    local.get 2
                    i32.load
                    local.get 3
                    i32.const 8
                    i32.add
                    local.get 2
                    i32.const 4
                    i32.add
                    i32.load
                    call_indirect (type 4)
                    br_if 5 (;@3;)
                    local.get 6
                    local.get 10
                    i32.ge_u
                    br_if 4 (;@4;)
                    local.get 5
                    i32.const -4
                    i32.add
                    local.set 2
                    local.get 5
                    i32.load
                    local.set 4
                    local.get 5
                    i32.const 8
                    i32.add
                    local.set 5
                    local.get 7
                    i32.const 36
                    i32.add
                    local.set 7
                    local.get 6
                    i32.const 1
                    i32.add
                    local.set 6
                    local.get 3
                    i32.const 8
                    i32.add
                    i32.const 24
                    i32.add
                    i32.load
                    local.get 2
                    i32.load
                    local.get 4
                    local.get 3
                    i32.const 8
                    i32.add
                    i32.const 28
                    i32.add
                    i32.load
                    i32.load offset=12
                    call_indirect (type 3)
                    i32.eqz
                    br_if 1 (;@7;)
                    br 5 (;@3;)
                  end
                end
                i32.const 1051144
                local.get 2
                local.get 4
                call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
                unreachable
              end
              i32.const 1050372
              call $_ZN4core9panicking5panic17h62fdcfa056e70982E
              unreachable
            end
            local.get 2
            i32.load
            local.set 8
            local.get 2
            i32.load offset=4
            local.tee 9
            local.get 4
            local.get 4
            local.get 9
            i32.gt_u
            select
            local.tee 10
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            local.get 8
            i32.load
            local.get 8
            i32.load offset=4
            local.get 1
            i32.load offset=12
            call_indirect (type 3)
            br_if 1 (;@3;)
            local.get 8
            i32.const 12
            i32.add
            local.set 7
            local.get 3
            i32.const 32
            i32.add
            local.set 0
            local.get 3
            i32.const 36
            i32.add
            local.set 1
            i32.const 1
            local.set 6
            loop  ;; label = @5
              local.get 5
              i32.load
              local.get 3
              i32.const 8
              i32.add
              local.get 5
              i32.const 4
              i32.add
              i32.load
              call_indirect (type 4)
              br_if 2 (;@3;)
              local.get 6
              local.get 10
              i32.ge_u
              br_if 1 (;@4;)
              local.get 7
              i32.const -4
              i32.add
              local.set 2
              local.get 7
              i32.load
              local.set 4
              local.get 7
              i32.const 8
              i32.add
              local.set 7
              local.get 5
              i32.const 8
              i32.add
              local.set 5
              local.get 6
              i32.const 1
              i32.add
              local.set 6
              local.get 0
              i32.load
              local.get 2
              i32.load
              local.get 4
              local.get 1
              i32.load
              i32.load offset=12
              call_indirect (type 3)
              i32.eqz
              br_if 0 (;@5;)
              br 2 (;@3;)
            end
          end
          local.get 9
          local.get 6
          i32.le_u
          br_if 1 (;@2;)
          local.get 3
          i32.const 32
          i32.add
          i32.load
          local.get 8
          local.get 6
          i32.const 3
          i32.shl
          i32.add
          local.tee 7
          i32.load
          local.get 7
          i32.load offset=4
          local.get 3
          i32.const 36
          i32.add
          i32.load
          i32.load offset=12
          call_indirect (type 3)
          i32.eqz
          br_if 1 (;@2;)
        end
        i32.const 1
        local.set 7
        br 1 (;@1;)
      end
      i32.const 0
      local.set 7
    end
    local.get 3
    i32.const 64
    i32.add
    global.set 0
    local.get 7)
  (func $_ZN71_$LT$core..ops..range..Range$LT$Idx$GT$$u20$as$u20$core..fmt..Debug$GT$3fmt17h6105e754ea0aeafbE (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        local.get 1
        call $_ZN4core3fmt3num52_$LT$impl$u20$core..fmt..Debug$u20$for$u20$usize$GT$3fmt17h5c9a921f5b6ae7beE
        br_if 0 (;@2;)
        local.get 1
        i32.const 28
        i32.add
        i32.load
        local.set 3
        local.get 1
        i32.load offset=24
        local.set 4
        local.get 2
        i32.const 28
        i32.add
        i32.const 0
        i32.store
        local.get 2
        i32.const 1050184
        i32.store offset=24
        local.get 2
        i64.const 1
        i64.store offset=12 align=4
        local.get 2
        i32.const 1050188
        i32.store offset=8
        local.get 4
        local.get 3
        local.get 2
        i32.const 8
        i32.add
        call $_ZN4core3fmt5write17h8cfd01c67a4a46c9E
        i32.eqz
        br_if 1 (;@1;)
      end
      local.get 2
      i32.const 32
      i32.add
      global.set 0
      i32.const 1
      return
    end
    local.get 0
    i32.const 4
    i32.add
    local.get 1
    call $_ZN4core3fmt3num52_$LT$impl$u20$core..fmt..Debug$u20$for$u20$usize$GT$3fmt17h5c9a921f5b6ae7beE
    local.set 1
    local.get 2
    i32.const 32
    i32.add
    global.set 0
    local.get 1)
  (func $_ZN4core3fmt3num52_$LT$impl$u20$core..fmt..Debug$u20$for$u20$usize$GT$3fmt17h5c9a921f5b6ae7beE (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 128
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              i32.load
              local.tee 3
              i32.const 16
              i32.and
              br_if 0 (;@5;)
              local.get 0
              i32.load
              local.set 4
              local.get 3
              i32.const 32
              i32.and
              br_if 1 (;@4;)
              local.get 4
              i64.extend_i32_u
              i32.const 1
              local.get 1
              call $_ZN4core3fmt3num3imp7fmt_u6417h4648b2300e699242E
              local.set 0
              br 2 (;@3;)
            end
            local.get 0
            i32.load
            local.set 4
            i32.const 0
            local.set 0
            loop  ;; label = @5
              local.get 2
              local.get 0
              i32.add
              i32.const 127
              i32.add
              local.get 4
              i32.const 15
              i32.and
              local.tee 3
              i32.const 48
              i32.or
              local.get 3
              i32.const 87
              i32.add
              local.get 3
              i32.const 10
              i32.lt_u
              select
              i32.store8
              local.get 0
              i32.const -1
              i32.add
              local.set 0
              local.get 4
              i32.const 4
              i32.shr_u
              local.tee 4
              br_if 0 (;@5;)
            end
            local.get 0
            i32.const 128
            i32.add
            local.tee 4
            i32.const 129
            i32.ge_u
            br_if 2 (;@2;)
            local.get 1
            i32.const 1
            i32.const 1050852
            i32.const 2
            local.get 2
            local.get 0
            i32.add
            i32.const 128
            i32.add
            i32.const 0
            local.get 0
            i32.sub
            call $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE
            local.set 0
            br 1 (;@3;)
          end
          i32.const 0
          local.set 0
          loop  ;; label = @4
            local.get 2
            local.get 0
            i32.add
            i32.const 127
            i32.add
            local.get 4
            i32.const 15
            i32.and
            local.tee 3
            i32.const 48
            i32.or
            local.get 3
            i32.const 55
            i32.add
            local.get 3
            i32.const 10
            i32.lt_u
            select
            i32.store8
            local.get 0
            i32.const -1
            i32.add
            local.set 0
            local.get 4
            i32.const 4
            i32.shr_u
            local.tee 4
            br_if 0 (;@4;)
          end
          local.get 0
          i32.const 128
          i32.add
          local.tee 4
          i32.const 129
          i32.ge_u
          br_if 2 (;@1;)
          local.get 1
          i32.const 1
          i32.const 1050852
          i32.const 2
          local.get 2
          local.get 0
          i32.add
          i32.const 128
          i32.add
          i32.const 0
          local.get 0
          i32.sub
          call $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE
          local.set 0
        end
        local.get 2
        i32.const 128
        i32.add
        global.set 0
        local.get 0
        return
      end
      local.get 4
      i32.const 128
      call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
      unreachable
    end
    local.get 4
    i32.const 128
    call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
    unreachable)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h5fb1d47f0acdabccE (type 7) (param i32) (result i64)
    i64.const -670765639137414048)
  (func $_ZN60_$LT$core..cell..BorrowError$u20$as$u20$core..fmt..Debug$GT$3fmt17h1bceae61083d85a2E (type 4) (param i32 i32) (result i32)
    local.get 1
    i32.load offset=24
    i32.const 1050196
    i32.const 11
    local.get 1
    i32.const 28
    i32.add
    i32.load
    i32.load offset=12
    call_indirect (type 3))
  (func $_ZN63_$LT$core..cell..BorrowMutError$u20$as$u20$core..fmt..Debug$GT$3fmt17h94d425ffcd3a62fcE (type 4) (param i32 i32) (result i32)
    local.get 1
    i32.load offset=24
    i32.const 1050207
    i32.const 14
    local.get 1
    i32.const 28
    i32.add
    i32.load
    i32.load offset=12
    call_indirect (type 3))
  (func $_ZN4core5panic9PanicInfo7message17hbe01febe97d5f595E (type 2) (param i32) (result i32)
    local.get 0
    i32.load offset=8)
  (func $_ZN4core5panic9PanicInfo8location17h959e627e2fd51988E (type 2) (param i32) (result i32)
    local.get 0
    i32.const 12
    i32.add)
  (func $_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17hd813a4fb6bc448f6E (type 4) (param i32 i32) (result i32)
    local.get 1
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    call $_ZN4core3fmt9Formatter3pad17hacaebd5abd28adf1E)
  (func $_ZN4core5panic8Location20internal_constructor17h55a9de981ef3aeb7E (type 11) (param i32 i32 i32 i32 i32)
    local.get 0
    local.get 4
    i32.store offset=12
    local.get 0
    local.get 3
    i32.store offset=8
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func $_ZN4core5panic8Location4file17h9e415b81e07881c2E (type 1) (param i32 i32)
    local.get 0
    local.get 1
    i64.load align=4
    i64.store align=4)
  (func $_ZN4core5panic8Location4line17h4caf5525f6ea07a3E (type 2) (param i32) (result i32)
    local.get 0
    i32.load offset=8)
  (func $_ZN4core5panic8Location6column17h3de733c09fa199ffE (type 2) (param i32) (result i32)
    local.get 0
    i32.load offset=12)
  (func $_ZN4core5slice6memchr6memchr17h6bde110ce4c09380E (type 9) (param i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    i32.const 0
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 3
        i32.and
        local.tee 5
        i32.eqz
        br_if 0 (;@2;)
        i32.const 4
        local.get 5
        i32.sub
        local.tee 5
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        local.get 3
        local.get 5
        local.get 5
        local.get 3
        i32.gt_u
        select
        local.tee 4
        i32.add
        local.set 6
        i32.const 0
        local.set 5
        local.get 1
        i32.const 255
        i32.and
        local.set 7
        local.get 4
        local.set 8
        local.get 2
        local.set 9
        block  ;; label = @3
          block  ;; label = @4
            loop  ;; label = @5
              local.get 6
              local.get 9
              i32.sub
              i32.const 3
              i32.le_u
              br_if 1 (;@4;)
              local.get 5
              local.get 9
              i32.load8_u
              local.tee 10
              local.get 7
              i32.ne
              i32.add
              local.set 5
              local.get 10
              local.get 7
              i32.eq
              br_if 2 (;@3;)
              local.get 5
              local.get 9
              i32.const 1
              i32.add
              i32.load8_u
              local.tee 10
              local.get 7
              i32.ne
              i32.add
              local.set 5
              local.get 10
              local.get 7
              i32.eq
              br_if 2 (;@3;)
              local.get 5
              local.get 9
              i32.const 2
              i32.add
              i32.load8_u
              local.tee 10
              local.get 7
              i32.ne
              i32.add
              local.set 5
              local.get 10
              local.get 7
              i32.eq
              br_if 2 (;@3;)
              local.get 5
              local.get 9
              i32.const 3
              i32.add
              i32.load8_u
              local.tee 10
              local.get 7
              i32.ne
              i32.add
              local.set 5
              local.get 8
              i32.const -4
              i32.add
              local.set 8
              local.get 9
              i32.const 4
              i32.add
              local.set 9
              local.get 10
              local.get 7
              i32.ne
              br_if 0 (;@5;)
              br 2 (;@3;)
            end
          end
          i32.const 0
          local.set 7
          local.get 1
          i32.const 255
          i32.and
          local.set 6
          loop  ;; label = @4
            local.get 8
            i32.eqz
            br_if 2 (;@2;)
            local.get 9
            local.get 7
            i32.add
            local.set 10
            local.get 8
            i32.const -1
            i32.add
            local.set 8
            local.get 7
            i32.const 1
            i32.add
            local.set 7
            local.get 10
            i32.load8_u
            local.tee 10
            local.get 6
            i32.ne
            br_if 0 (;@4;)
          end
          local.get 10
          local.get 1
          i32.const 255
          i32.and
          i32.eq
          i32.const 1
          i32.add
          i32.const 1
          i32.and
          local.get 5
          i32.add
          local.get 7
          i32.add
          i32.const -1
          i32.add
          local.set 5
        end
        i32.const 1
        local.set 9
        br 1 (;@1;)
      end
      local.get 1
      i32.const 255
      i32.and
      local.set 7
      block  ;; label = @2
        block  ;; label = @3
          local.get 3
          i32.const 8
          i32.lt_u
          br_if 0 (;@3;)
          local.get 4
          local.get 3
          i32.const -8
          i32.add
          local.tee 10
          i32.gt_u
          br_if 0 (;@3;)
          local.get 7
          i32.const 16843009
          i32.mul
          local.set 5
          block  ;; label = @4
            loop  ;; label = @5
              local.get 2
              local.get 4
              i32.add
              local.tee 9
              i32.const 4
              i32.add
              i32.load
              local.get 5
              i32.xor
              local.tee 8
              i32.const -1
              i32.xor
              local.get 8
              i32.const -16843009
              i32.add
              i32.and
              local.get 9
              i32.load
              local.get 5
              i32.xor
              local.tee 9
              i32.const -1
              i32.xor
              local.get 9
              i32.const -16843009
              i32.add
              i32.and
              i32.or
              i32.const -2139062144
              i32.and
              br_if 1 (;@4;)
              local.get 4
              i32.const 8
              i32.add
              local.tee 4
              local.get 10
              i32.le_u
              br_if 0 (;@5;)
            end
          end
          local.get 4
          local.get 3
          i32.gt_u
          br_if 1 (;@2;)
        end
        local.get 2
        local.get 4
        i32.add
        local.set 9
        local.get 2
        local.get 3
        i32.add
        local.set 2
        local.get 3
        local.get 4
        i32.sub
        local.set 8
        i32.const 0
        local.set 5
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              loop  ;; label = @6
                local.get 2
                local.get 9
                i32.sub
                i32.const 3
                i32.le_u
                br_if 1 (;@5;)
                local.get 5
                local.get 9
                i32.load8_u
                local.tee 10
                local.get 7
                i32.ne
                i32.add
                local.set 5
                local.get 10
                local.get 7
                i32.eq
                br_if 2 (;@4;)
                local.get 5
                local.get 9
                i32.const 1
                i32.add
                i32.load8_u
                local.tee 10
                local.get 7
                i32.ne
                i32.add
                local.set 5
                local.get 10
                local.get 7
                i32.eq
                br_if 2 (;@4;)
                local.get 5
                local.get 9
                i32.const 2
                i32.add
                i32.load8_u
                local.tee 10
                local.get 7
                i32.ne
                i32.add
                local.set 5
                local.get 10
                local.get 7
                i32.eq
                br_if 2 (;@4;)
                local.get 5
                local.get 9
                i32.const 3
                i32.add
                i32.load8_u
                local.tee 10
                local.get 7
                i32.ne
                i32.add
                local.set 5
                local.get 8
                i32.const -4
                i32.add
                local.set 8
                local.get 9
                i32.const 4
                i32.add
                local.set 9
                local.get 10
                local.get 7
                i32.ne
                br_if 0 (;@6;)
                br 2 (;@4;)
              end
            end
            i32.const 0
            local.set 7
            local.get 1
            i32.const 255
            i32.and
            local.set 2
            loop  ;; label = @5
              local.get 8
              i32.eqz
              br_if 2 (;@3;)
              local.get 9
              local.get 7
              i32.add
              local.set 10
              local.get 8
              i32.const -1
              i32.add
              local.set 8
              local.get 7
              i32.const 1
              i32.add
              local.set 7
              local.get 10
              i32.load8_u
              local.tee 10
              local.get 2
              i32.ne
              br_if 0 (;@5;)
            end
            local.get 10
            local.get 1
            i32.const 255
            i32.and
            i32.eq
            i32.const 1
            i32.add
            i32.const 1
            i32.and
            local.get 5
            i32.add
            local.get 7
            i32.add
            i32.const -1
            i32.add
            local.set 5
          end
          i32.const 1
          local.set 9
          local.get 5
          local.get 4
          i32.add
          local.set 5
          br 2 (;@1;)
        end
        i32.const 0
        local.set 9
        local.get 5
        local.get 7
        i32.add
        local.get 4
        i32.add
        local.set 5
        br 1 (;@1;)
      end
      local.get 4
      local.get 3
      call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
      unreachable
    end
    local.get 0
    local.get 5
    i32.store offset=4
    local.get 0
    local.get 9
    i32.store)
  (func $_ZN4core7unicode9bool_trie8BoolTrie6lookup17h03e59958127816b1E (type 4) (param i32 i32) (result i32)
    (local i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i32.const 2047
                  i32.gt_u
                  br_if 0 (;@7;)
                  local.get 0
                  local.get 1
                  i32.const 3
                  i32.shr_u
                  i32.const 536870904
                  i32.and
                  i32.add
                  local.set 0
                  br 1 (;@6;)
                end
                block  ;; label = @7
                  local.get 1
                  i32.const 65535
                  i32.gt_u
                  br_if 0 (;@7;)
                  local.get 1
                  i32.const 6
                  i32.shr_u
                  i32.const -32
                  i32.add
                  local.tee 2
                  i32.const 991
                  i32.gt_u
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const 260
                  i32.add
                  i32.load
                  local.tee 3
                  local.get 0
                  local.get 2
                  i32.add
                  i32.const 280
                  i32.add
                  i32.load8_u
                  local.tee 2
                  i32.le_u
                  br_if 3 (;@4;)
                  local.get 0
                  i32.load offset=256
                  local.get 2
                  i32.const 3
                  i32.shl
                  i32.add
                  local.set 0
                  br 1 (;@6;)
                end
                local.get 1
                i32.const 12
                i32.shr_u
                i32.const -16
                i32.add
                local.tee 2
                i32.const 255
                i32.gt_u
                br_if 3 (;@3;)
                local.get 0
                local.get 2
                i32.add
                i32.const 1272
                i32.add
                i32.load8_u
                i32.const 6
                i32.shl
                local.get 1
                i32.const 6
                i32.shr_u
                i32.const 63
                i32.and
                i32.or
                local.tee 2
                local.get 0
                i32.const 268
                i32.add
                i32.load
                local.tee 3
                i32.ge_u
                br_if 4 (;@2;)
                local.get 0
                i32.const 276
                i32.add
                i32.load
                local.tee 3
                local.get 0
                i32.load offset=264
                local.get 2
                i32.add
                i32.load8_u
                local.tee 2
                i32.le_u
                br_if 5 (;@1;)
                local.get 0
                i32.load offset=272
                local.get 2
                i32.const 3
                i32.shl
                i32.add
                local.set 0
              end
              local.get 0
              i64.load
              i64.const 1
              local.get 1
              i32.const 63
              i32.and
              i64.extend_i32_u
              i64.shl
              i64.and
              i64.const 0
              i64.ne
              return
            end
            i32.const 1051216
            local.get 2
            i32.const 992
            call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
            unreachable
          end
          i32.const 1051232
          local.get 2
          local.get 3
          call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
          unreachable
        end
        i32.const 1051248
        local.get 2
        i32.const 256
        call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
        unreachable
      end
      i32.const 1051264
      local.get 2
      local.get 3
      call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
      unreachable
    end
    i32.const 1051280
    local.get 2
    local.get 3
    call $_ZN4core9panicking18panic_bounds_check17h51667a9f831439a7E
    unreachable)
  (func $_ZN4core7unicode9printable5check17hc99709d493feb50bE (type 12) (param i32 i32 i32 i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    i32.const 1
    local.set 7
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 2
                i32.eqz
                br_if 0 (;@6;)
                local.get 1
                local.get 2
                i32.const 1
                i32.shl
                i32.add
                local.set 8
                local.get 0
                i32.const 65280
                i32.and
                i32.const 8
                i32.shr_u
                local.set 9
                i32.const 0
                local.set 10
                local.get 0
                i32.const 255
                i32.and
                local.set 11
                loop  ;; label = @7
                  local.get 1
                  i32.const 2
                  i32.add
                  local.set 12
                  local.get 10
                  local.get 1
                  i32.load8_u offset=1
                  local.tee 2
                  i32.add
                  local.set 13
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 1
                      i32.load8_u
                      local.tee 1
                      local.get 9
                      i32.ne
                      br_if 0 (;@9;)
                      local.get 13
                      local.get 10
                      i32.lt_u
                      br_if 7 (;@2;)
                      local.get 13
                      local.get 4
                      i32.gt_u
                      br_if 8 (;@1;)
                      local.get 3
                      local.get 10
                      i32.add
                      local.set 1
                      loop  ;; label = @10
                        local.get 2
                        i32.eqz
                        br_if 2 (;@8;)
                        local.get 2
                        i32.const -1
                        i32.add
                        local.set 2
                        local.get 1
                        i32.load8_u
                        local.set 10
                        local.get 1
                        i32.const 1
                        i32.add
                        local.set 1
                        local.get 10
                        local.get 11
                        i32.ne
                        br_if 0 (;@10;)
                        br 5 (;@5;)
                      end
                    end
                    local.get 1
                    local.get 9
                    i32.gt_u
                    br_if 2 (;@6;)
                    local.get 13
                    local.set 10
                    local.get 12
                    local.set 1
                    local.get 12
                    local.get 8
                    i32.ne
                    br_if 1 (;@7;)
                    br 2 (;@6;)
                  end
                  local.get 13
                  local.set 10
                  local.get 12
                  local.set 1
                  local.get 12
                  local.get 8
                  i32.ne
                  br_if 0 (;@7;)
                end
              end
              local.get 6
              i32.eqz
              br_if 1 (;@4;)
              local.get 5
              local.get 6
              i32.add
              local.set 11
              local.get 0
              i32.const 65535
              i32.and
              local.set 1
              i32.const 1
              local.set 7
              loop  ;; label = @6
                local.get 5
                i32.const 1
                i32.add
                local.set 10
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 5
                    i32.load8_u
                    local.tee 2
                    i32.const 24
                    i32.shl
                    i32.const 24
                    i32.shr_s
                    local.tee 13
                    i32.const -1
                    i32.le_s
                    br_if 0 (;@8;)
                    local.get 10
                    local.set 5
                    br 1 (;@7;)
                  end
                  local.get 10
                  local.get 11
                  i32.eq
                  br_if 4 (;@3;)
                  local.get 13
                  i32.const 127
                  i32.and
                  i32.const 8
                  i32.shl
                  local.get 5
                  i32.const 1
                  i32.add
                  i32.load8_u
                  i32.or
                  local.set 2
                  local.get 5
                  i32.const 2
                  i32.add
                  local.set 5
                end
                local.get 1
                local.get 2
                i32.sub
                local.tee 1
                i32.const 0
                i32.lt_s
                br_if 2 (;@4;)
                local.get 7
                i32.const 1
                i32.xor
                local.set 7
                local.get 5
                local.get 11
                i32.ne
                br_if 0 (;@6;)
                br 2 (;@4;)
              end
            end
            i32.const 0
            local.set 7
          end
          local.get 7
          i32.const 1
          i32.and
          return
        end
        i32.const 1050372
        call $_ZN4core9panicking5panic17h62fdcfa056e70982E
        unreachable
      end
      local.get 10
      local.get 13
      call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
      unreachable
    end
    local.get 13
    local.get 4
    call $_ZN4core5slice20slice_index_len_fail17hb115deb2b20f49d8E
    unreachable)
  (func $_ZN4core3str6traits101_$LT$impl$u20$core..slice..SliceIndex$LT$str$GT$$u20$for$u20$core..ops..range..Range$LT$usize$GT$$GT$5index28_$u7b$$u7b$closure$u7d$$u7d$17h5fc12def9a7684b5E (type 0) (param i32)
    (local i32)
    local.get 0
    i32.load
    local.tee 1
    i32.load
    local.get 1
    i32.load offset=4
    local.get 0
    i32.load offset=4
    i32.load
    local.get 0
    i32.load offset=8
    i32.load
    call $_ZN4core3str16slice_error_fail17h3704ce74b976be71E
    unreachable)
  (func $_ZN4core3fmt3num52_$LT$impl$u20$core..fmt..LowerHex$u20$for$u20$i8$GT$3fmt17hac36f7f0417b1c1dE (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 128
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    i32.load8_u
    local.set 3
    i32.const 0
    local.set 0
    loop  ;; label = @1
      local.get 2
      local.get 0
      i32.add
      i32.const 127
      i32.add
      local.get 3
      i32.const 15
      i32.and
      local.tee 4
      i32.const 48
      i32.or
      local.get 4
      i32.const 87
      i32.add
      local.get 4
      i32.const 10
      i32.lt_u
      select
      i32.store8
      local.get 0
      i32.const -1
      i32.add
      local.set 0
      local.get 3
      i32.const 4
      i32.shr_u
      i32.const 15
      i32.and
      local.tee 3
      br_if 0 (;@1;)
    end
    block  ;; label = @1
      local.get 0
      i32.const 128
      i32.add
      local.tee 3
      i32.const 129
      i32.ge_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 1
      i32.const 1050852
      i32.const 2
      local.get 2
      local.get 0
      i32.add
      i32.const 128
      i32.add
      i32.const 0
      local.get 0
      i32.sub
      call $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE
      local.set 0
      local.get 2
      i32.const 128
      i32.add
      global.set 0
      local.get 0
      return
    end
    local.get 3
    i32.const 128
    call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
    unreachable)
  (func $_ZN4core3fmt3num3imp51_$LT$impl$u20$core..fmt..Display$u20$for$u20$u8$GT$3fmt17ha887097ae9f6ec79E (type 4) (param i32 i32) (result i32)
    local.get 0
    i64.load8_u
    i32.const 1
    local.get 1
    call $_ZN4core3fmt3num3imp7fmt_u6417h4648b2300e699242E)
  (func $_ZN41_$LT$char$u20$as$u20$core..fmt..Debug$GT$3fmt17h7a420490f85a5bf3E (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32 i64 i32 i32)
    i32.const 1
    local.set 2
    block  ;; label = @1
      local.get 1
      i32.load offset=24
      i32.const 39
      local.get 1
      i32.const 28
      i32.add
      i32.load
      i32.load offset=16
      call_indirect (type 4)
      br_if 0 (;@1;)
      i32.const 2
      local.set 2
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 0
                          i32.load
                          local.tee 0
                          i32.const -9
                          i32.add
                          local.tee 3
                          i32.const 30
                          i32.gt_u
                          br_if 0 (;@11;)
                          i32.const 116
                          local.set 4
                          block  ;; label = @12
                            local.get 3
                            br_table 10 (;@2;) 0 (;@12;) 2 (;@10;) 2 (;@10;) 3 (;@9;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 6 (;@6;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 2 (;@10;) 6 (;@6;) 10 (;@2;)
                          end
                          i32.const 110
                          local.set 4
                          br 3 (;@8;)
                        end
                        local.get 0
                        i32.const 92
                        i32.eq
                        br_if 4 (;@6;)
                      end
                      i32.const 1054144
                      local.get 0
                      call $_ZN4core7unicode9bool_trie8BoolTrie6lookup17h03e59958127816b1E
                      i32.eqz
                      br_if 2 (;@7;)
                      local.get 0
                      i32.const 1
                      i32.or
                      i32.clz
                      i32.const 2
                      i32.shr_u
                      i32.const 7
                      i32.xor
                      i64.extend_i32_u
                      i64.const 21474836480
                      i64.or
                      local.set 5
                      br 5 (;@4;)
                    end
                    i32.const 114
                    local.set 4
                  end
                  br 5 (;@2;)
                end
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 0
                    i32.const 65535
                    i32.gt_u
                    br_if 0 (;@8;)
                    local.get 0
                    i32.const 1051296
                    i32.const 40
                    i32.const 1051376
                    i32.const 303
                    i32.const 1051679
                    i32.const 316
                    call $_ZN4core7unicode9printable5check17hc99709d493feb50bE
                    i32.eqz
                    br_if 3 (;@5;)
                    br 1 (;@7;)
                  end
                  block  ;; label = @8
                    local.get 0
                    i32.const 131071
                    i32.gt_u
                    br_if 0 (;@8;)
                    local.get 0
                    i32.const 1051995
                    i32.const 33
                    i32.const 1052061
                    i32.const 158
                    i32.const 1052219
                    i32.const 381
                    call $_ZN4core7unicode9printable5check17hc99709d493feb50bE
                    br_if 1 (;@7;)
                    br 3 (;@5;)
                  end
                  local.get 0
                  i32.const 917999
                  i32.gt_u
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const -195102
                  i32.add
                  i32.const 722658
                  i32.lt_u
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const -191457
                  i32.add
                  i32.const 3103
                  i32.lt_u
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const -183970
                  i32.add
                  i32.const 14
                  i32.lt_u
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const 2097150
                  i32.and
                  i32.const 178206
                  i32.eq
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const -173783
                  i32.add
                  i32.const 41
                  i32.lt_u
                  br_if 2 (;@5;)
                  local.get 0
                  i32.const -177973
                  i32.add
                  i32.const 10
                  i32.le_u
                  br_if 2 (;@5;)
                end
                i32.const 1
                local.set 2
              end
              br 2 (;@3;)
            end
            local.get 0
            i32.const 1
            i32.or
            i32.clz
            i32.const 2
            i32.shr_u
            i32.const 7
            i32.xor
            i64.extend_i32_u
            i64.const 21474836480
            i64.or
            local.set 5
          end
          i32.const 3
          local.set 2
        end
        local.get 0
        local.set 4
      end
      local.get 1
      i32.const 24
      i32.add
      local.set 3
      local.get 1
      i32.const 28
      i32.add
      local.set 6
      block  ;; label = @2
        loop  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            local.get 2
                            i32.const 1
                            i32.eq
                            br_if 0 (;@12;)
                            i32.const 92
                            local.set 0
                            local.get 2
                            i32.const 2
                            i32.eq
                            br_if 1 (;@11;)
                            local.get 2
                            i32.const 3
                            i32.ne
                            br_if 10 (;@2;)
                            local.get 5
                            i64.const 32
                            i64.shr_u
                            i32.wrap_i64
                            i32.const 255
                            i32.and
                            i32.const -1
                            i32.add
                            local.tee 2
                            i32.const 4
                            i32.gt_u
                            br_if 10 (;@2;)
                            block  ;; label = @13
                              local.get 2
                              br_table 0 (;@13;) 3 (;@10;) 4 (;@9;) 5 (;@8;) 6 (;@7;) 0 (;@13;)
                            end
                            local.get 5
                            i64.const -1095216660481
                            i64.and
                            local.set 5
                            i32.const 125
                            local.set 0
                            br 7 (;@5;)
                          end
                          i32.const 0
                          local.set 2
                          local.get 4
                          local.set 0
                          br 7 (;@4;)
                        end
                        i32.const 1
                        local.set 2
                        br 6 (;@4;)
                      end
                      local.get 4
                      local.get 5
                      i32.wrap_i64
                      local.tee 7
                      i32.const 2
                      i32.shl
                      i32.const 28
                      i32.and
                      i32.shr_u
                      i32.const 15
                      i32.and
                      local.tee 2
                      i32.const 48
                      i32.or
                      local.get 2
                      i32.const 87
                      i32.add
                      local.get 2
                      i32.const 10
                      i32.lt_u
                      select
                      local.set 0
                      local.get 7
                      i32.eqz
                      br_if 3 (;@6;)
                      local.get 5
                      i64.const -1
                      i64.add
                      i64.const 4294967295
                      i64.and
                      local.get 5
                      i64.const -4294967296
                      i64.and
                      i64.or
                      local.set 5
                      br 4 (;@5;)
                    end
                    local.get 5
                    i64.const -1095216660481
                    i64.and
                    i64.const 8589934592
                    i64.or
                    local.set 5
                    i32.const 123
                    local.set 0
                    br 3 (;@5;)
                  end
                  local.get 5
                  i64.const -1095216660481
                  i64.and
                  i64.const 12884901888
                  i64.or
                  local.set 5
                  i32.const 117
                  local.set 0
                  br 2 (;@5;)
                end
                local.get 5
                i64.const -1095216660481
                i64.and
                i64.const 17179869184
                i64.or
                local.set 5
                br 1 (;@5;)
              end
              local.get 5
              i64.const -1095216660481
              i64.and
              i64.const 4294967296
              i64.or
              local.set 5
            end
            i32.const 3
            local.set 2
          end
          local.get 3
          i32.load
          local.get 0
          local.get 6
          i32.load
          i32.load offset=16
          call_indirect (type 4)
          i32.eqz
          br_if 0 (;@3;)
        end
        i32.const 1
        return
      end
      local.get 1
      i32.const 24
      i32.add
      i32.load
      i32.const 39
      local.get 1
      i32.const 28
      i32.add
      i32.load
      i32.load offset=16
      call_indirect (type 4)
      local.set 2
    end
    local.get 2)
  (func $_ZN68_$LT$core..fmt..builders..PadAdapter$u20$as$u20$core..fmt..Write$GT$9write_str17h7d69c319e62567eaE (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 3
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.eqz
              br_if 0 (;@5;)
              local.get 3
              i32.const 40
              i32.add
              local.set 4
              local.get 0
              i32.const 8
              i32.add
              local.set 5
              local.get 3
              i32.const 32
              i32.add
              local.set 6
              local.get 3
              i32.const 28
              i32.add
              local.set 7
              local.get 3
              i32.const 36
              i32.add
              local.set 8
              local.get 0
              i32.const 4
              i32.add
              local.set 9
              loop  ;; label = @6
                block  ;; label = @7
                  local.get 5
                  i32.load8_u
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 0
                  i32.load
                  i32.const 1051080
                  i32.const 4
                  local.get 9
                  i32.load
                  i32.load offset=12
                  call_indirect (type 3)
                  br_if 3 (;@4;)
                end
                local.get 4
                i32.const 10
                i32.store
                local.get 6
                i64.const 4294967306
                i64.store
                local.get 7
                local.get 2
                i32.store
                local.get 3
                i32.const 16
                i32.add
                i32.const 8
                i32.add
                local.tee 10
                i32.const 0
                i32.store
                local.get 3
                local.get 2
                i32.store offset=20
                local.get 3
                local.get 1
                i32.store offset=16
                local.get 3
                i32.const 8
                i32.add
                i32.const 10
                local.get 1
                local.get 2
                call $_ZN4core5slice6memchr6memchr17h6bde110ce4c09380E
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 3
                          i32.load offset=8
                          i32.const 1
                          i32.ne
                          br_if 0 (;@11;)
                          local.get 3
                          i32.load offset=12
                          local.set 11
                          loop  ;; label = @12
                            local.get 10
                            local.get 11
                            local.get 10
                            i32.load
                            i32.add
                            i32.const 1
                            i32.add
                            local.tee 11
                            i32.store
                            block  ;; label = @13
                              block  ;; label = @14
                                local.get 11
                                local.get 8
                                i32.load
                                local.tee 12
                                i32.ge_u
                                br_if 0 (;@14;)
                                local.get 3
                                i32.load offset=20
                                local.set 13
                                br 1 (;@13;)
                              end
                              local.get 3
                              i32.load offset=20
                              local.tee 13
                              local.get 11
                              i32.lt_u
                              br_if 0 (;@13;)
                              local.get 12
                              i32.const 5
                              i32.ge_u
                              br_if 5 (;@8;)
                              local.get 3
                              i32.load offset=16
                              local.get 11
                              local.get 12
                              i32.sub
                              local.tee 14
                              i32.add
                              local.tee 15
                              local.get 4
                              i32.eq
                              br_if 4 (;@9;)
                              local.get 15
                              local.get 4
                              local.get 12
                              call $memcmp
                              i32.eqz
                              br_if 4 (;@9;)
                            end
                            local.get 7
                            i32.load
                            local.tee 15
                            local.get 11
                            i32.lt_u
                            br_if 2 (;@10;)
                            local.get 13
                            local.get 15
                            i32.lt_u
                            br_if 2 (;@10;)
                            local.get 3
                            local.get 3
                            i32.const 16
                            i32.add
                            local.get 12
                            i32.add
                            i32.const 23
                            i32.add
                            i32.load8_u
                            local.get 3
                            i32.load offset=16
                            local.get 11
                            i32.add
                            local.get 15
                            local.get 11
                            i32.sub
                            call $_ZN4core5slice6memchr6memchr17h6bde110ce4c09380E
                            local.get 3
                            i32.load offset=4
                            local.set 11
                            local.get 3
                            i32.load
                            i32.const 1
                            i32.eq
                            br_if 0 (;@12;)
                          end
                        end
                        local.get 10
                        local.get 7
                        i32.load
                        i32.store
                      end
                      local.get 5
                      i32.const 0
                      i32.store8
                      local.get 2
                      local.set 11
                      br 2 (;@7;)
                    end
                    local.get 5
                    i32.const 1
                    i32.store8
                    local.get 14
                    i32.const 1
                    i32.add
                    local.set 11
                    br 1 (;@7;)
                  end
                  local.get 12
                  i32.const 4
                  call $_ZN4core5slice20slice_index_len_fail17hb115deb2b20f49d8E
                  unreachable
                end
                local.get 9
                i32.load
                local.set 15
                local.get 0
                i32.load
                local.set 12
                block  ;; label = @7
                  local.get 11
                  i32.eqz
                  local.get 2
                  local.get 11
                  i32.eq
                  i32.or
                  local.tee 10
                  br_if 0 (;@7;)
                  local.get 2
                  local.get 11
                  i32.le_u
                  br_if 5 (;@2;)
                  local.get 1
                  local.get 11
                  i32.add
                  i32.load8_s
                  i32.const -65
                  i32.le_s
                  br_if 5 (;@2;)
                end
                local.get 12
                local.get 1
                local.get 11
                local.get 15
                i32.load offset=12
                call_indirect (type 3)
                br_if 2 (;@4;)
                block  ;; label = @7
                  local.get 10
                  br_if 0 (;@7;)
                  local.get 2
                  local.get 11
                  i32.le_u
                  br_if 6 (;@1;)
                  local.get 1
                  local.get 11
                  i32.add
                  i32.load8_s
                  i32.const -65
                  i32.le_s
                  br_if 6 (;@1;)
                end
                local.get 1
                local.get 11
                i32.add
                local.set 1
                local.get 2
                local.get 11
                i32.sub
                local.tee 2
                br_if 0 (;@6;)
              end
            end
            i32.const 0
            local.set 11
            br 1 (;@3;)
          end
          i32.const 1
          local.set 11
        end
        local.get 3
        i32.const 48
        i32.add
        global.set 0
        local.get 11
        return
      end
      local.get 1
      local.get 2
      i32.const 0
      local.get 11
      call $_ZN4core3str16slice_error_fail17h3704ce74b976be71E
      unreachable
    end
    local.get 1
    local.get 2
    local.get 11
    local.get 2
    call $_ZN4core3str16slice_error_fail17h3704ce74b976be71E
    unreachable)
  (func $_ZN4core3fmt8builders10DebugTuple5field17hc049fbe3a16238abE (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i64 i64 i64 i64 i64)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 3
    global.set 0
    i32.const 1
    local.set 4
    block  ;; label = @1
      local.get 0
      i32.load8_u offset=8
      br_if 0 (;@1;)
      local.get 0
      i32.load offset=4
      local.set 5
      block  ;; label = @2
        local.get 0
        i32.load
        local.tee 6
        i32.load8_u
        i32.const 4
        i32.and
        br_if 0 (;@2;)
        i32.const 1
        local.set 4
        local.get 6
        i32.load offset=24
        i32.const 1051086
        i32.const 1051090
        local.get 5
        select
        i32.const 2
        i32.const 1
        local.get 5
        select
        local.get 6
        i32.const 28
        i32.add
        i32.load
        i32.load offset=12
        call_indirect (type 3)
        br_if 1 (;@1;)
        local.get 1
        local.get 0
        i32.load
        local.get 2
        i32.load offset=12
        call_indirect (type 4)
        local.set 4
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 5
        br_if 0 (;@2;)
        i32.const 1
        local.set 4
        local.get 6
        i32.load offset=24
        i32.const 1051088
        i32.const 2
        local.get 6
        i32.const 28
        i32.add
        i32.load
        i32.load offset=12
        call_indirect (type 3)
        br_if 1 (;@1;)
        local.get 0
        i32.load
        local.set 6
      end
      i32.const 1
      local.set 4
      local.get 3
      i32.const 1
      i32.store8 offset=16
      local.get 6
      i64.load offset=16 align=4
      local.set 7
      local.get 6
      i64.load offset=8 align=4
      local.set 8
      local.get 3
      i32.const 52
      i32.add
      local.tee 5
      i32.const 1051056
      i32.store
      local.get 3
      local.get 6
      i64.load offset=24 align=4
      i64.store offset=8
      local.get 6
      i64.load offset=32 align=4
      local.set 9
      local.get 6
      i64.load offset=40 align=4
      local.set 10
      local.get 3
      local.get 6
      i32.load8_u offset=48
      i32.store8 offset=72
      local.get 6
      i64.load align=4
      local.set 11
      local.get 3
      local.get 8
      i64.store offset=32
      local.get 3
      local.get 7
      i64.store offset=40
      local.get 3
      local.get 10
      i64.store offset=64
      local.get 3
      local.get 9
      i64.store offset=56
      local.get 3
      local.get 11
      i64.store offset=24
      local.get 3
      local.get 3
      i32.const 8
      i32.add
      i32.store offset=48
      local.get 1
      local.get 3
      i32.const 24
      i32.add
      local.get 2
      i32.load offset=12
      call_indirect (type 4)
      br_if 0 (;@1;)
      local.get 3
      i32.const 48
      i32.add
      i32.load
      i32.const 1051084
      i32.const 2
      local.get 5
      i32.load
      i32.load offset=12
      call_indirect (type 3)
      local.set 4
    end
    local.get 0
    i32.const 8
    i32.add
    local.get 4
    i32.store8
    local.get 0
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.add
    i32.store offset=4
    local.get 3
    i32.const 80
    i32.add
    global.set 0
    local.get 0)
  (func $_ZN4core3fmt8builders10DebugTuple6finish17h9e783d267307e9eaE (type 2) (param i32) (result i32)
    (local i32 i32 i32)
    local.get 0
    i32.load8_u offset=8
    local.set 1
    block  ;; label = @1
      local.get 0
      i32.load offset=4
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 255
      i32.and
      local.set 3
      i32.const 1
      local.set 1
      block  ;; label = @2
        local.get 3
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 2
          i32.const 1
          i32.ne
          br_if 0 (;@3;)
          local.get 0
          i32.load8_u offset=9
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.tee 3
          i32.load8_u
          i32.const 4
          i32.and
          br_if 0 (;@3;)
          i32.const 1
          local.set 1
          local.get 3
          i32.load offset=24
          i32.const 1051091
          i32.const 1
          local.get 3
          i32.const 28
          i32.add
          i32.load
          i32.load offset=12
          call_indirect (type 3)
          br_if 1 (;@2;)
        end
        local.get 0
        i32.load
        local.tee 1
        i32.load offset=24
        i32.const 1051092
        i32.const 1
        local.get 1
        i32.const 28
        i32.add
        i32.load
        i32.load offset=12
        call_indirect (type 3)
        local.set 1
      end
      local.get 0
      i32.const 8
      i32.add
      local.get 1
      i32.store8
    end
    local.get 1
    i32.const 255
    i32.and
    i32.const 0
    i32.ne)
  (func $_ZN4core3fmt8builders10DebugInner5entry17hda957e0b29db8990E (type 8) (param i32 i32 i32)
    (local i32 i32 i32 i64 i64 i32 i64 i64 i64)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 3
    global.set 0
    i32.const 1
    local.set 4
    block  ;; label = @1
      local.get 0
      i32.load8_u offset=4
      br_if 0 (;@1;)
      local.get 0
      i32.load8_u offset=5
      local.set 4
      block  ;; label = @2
        local.get 0
        i32.load
        local.tee 5
        i32.load8_u
        i32.const 4
        i32.and
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 4
          i32.const 255
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          i32.const 1
          local.set 4
          local.get 5
          i32.load offset=24
          i32.const 1051086
          i32.const 2
          local.get 5
          i32.const 28
          i32.add
          i32.load
          i32.load offset=12
          call_indirect (type 3)
          br_if 2 (;@1;)
          local.get 0
          i32.load
          local.set 5
        end
        local.get 1
        local.get 5
        local.get 2
        i32.load offset=12
        call_indirect (type 4)
        local.set 4
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 4
        i32.const 255
        i32.and
        br_if 0 (;@2;)
        i32.const 1
        local.set 4
        local.get 5
        i32.load offset=24
        i32.const 1051093
        i32.const 1
        local.get 5
        i32.const 28
        i32.add
        i32.load
        i32.load offset=12
        call_indirect (type 3)
        br_if 1 (;@1;)
        local.get 0
        i32.load
        local.set 5
      end
      i32.const 1
      local.set 4
      local.get 3
      i32.const 1
      i32.store8 offset=16
      local.get 5
      i64.load offset=16 align=4
      local.set 6
      local.get 5
      i64.load offset=8 align=4
      local.set 7
      local.get 3
      i32.const 52
      i32.add
      local.tee 8
      i32.const 1051056
      i32.store
      local.get 3
      local.get 5
      i64.load offset=24 align=4
      i64.store offset=8
      local.get 5
      i64.load offset=32 align=4
      local.set 9
      local.get 5
      i64.load offset=40 align=4
      local.set 10
      local.get 3
      local.get 5
      i32.load8_u offset=48
      i32.store8 offset=72
      local.get 5
      i64.load align=4
      local.set 11
      local.get 3
      local.get 7
      i64.store offset=32
      local.get 3
      local.get 6
      i64.store offset=40
      local.get 3
      local.get 10
      i64.store offset=64
      local.get 3
      local.get 9
      i64.store offset=56
      local.get 3
      local.get 11
      i64.store offset=24
      local.get 3
      local.get 3
      i32.const 8
      i32.add
      i32.store offset=48
      local.get 1
      local.get 3
      i32.const 24
      i32.add
      local.get 2
      i32.load offset=12
      call_indirect (type 4)
      br_if 0 (;@1;)
      local.get 3
      i32.const 48
      i32.add
      i32.load
      i32.const 1051084
      i32.const 2
      local.get 8
      i32.load
      i32.load offset=12
      call_indirect (type 3)
      local.set 4
    end
    local.get 0
    i32.const 1
    i32.store8 offset=5
    local.get 0
    i32.const 4
    i32.add
    local.get 4
    i32.store8
    local.get 3
    i32.const 80
    i32.add
    global.set 0)
  (func $_ZN4core3fmt8builders8DebugSet5entry17hdfa2936ae55a6785E (type 3) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call $_ZN4core3fmt8builders10DebugInner5entry17hda957e0b29db8990E
    local.get 0)
  (func $_ZN4core3fmt8builders9DebugList6finish17hdbd1049d6264d19cE (type 2) (param i32) (result i32)
    block  ;; label = @1
      local.get 0
      i32.load8_u offset=4
      i32.eqz
      br_if 0 (;@1;)
      i32.const 1
      return
    end
    local.get 0
    i32.load
    local.tee 0
    i32.load offset=24
    i32.const 1051095
    i32.const 1
    local.get 0
    i32.const 28
    i32.add
    i32.load
    i32.load offset=12
    call_indirect (type 3))
  (func $_ZN4core3fmt5Write10write_char17he6255f3eabfceb48E (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 0
    i32.store offset=12
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 127
        i32.gt_u
        br_if 0 (;@2;)
        local.get 2
        local.get 1
        i32.store8 offset=12
        i32.const 1
        local.set 1
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 1
        i32.const 2047
        i32.gt_u
        br_if 0 (;@2;)
        local.get 2
        local.get 1
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=13
        local.get 2
        local.get 1
        i32.const 6
        i32.shr_u
        i32.const 31
        i32.and
        i32.const 192
        i32.or
        i32.store8 offset=12
        i32.const 2
        local.set 1
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 1
        i32.const 65535
        i32.gt_u
        br_if 0 (;@2;)
        local.get 2
        local.get 1
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=14
        local.get 2
        local.get 1
        i32.const 6
        i32.shr_u
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=13
        local.get 2
        local.get 1
        i32.const 12
        i32.shr_u
        i32.const 15
        i32.and
        i32.const 224
        i32.or
        i32.store8 offset=12
        i32.const 3
        local.set 1
        br 1 (;@1;)
      end
      local.get 2
      local.get 1
      i32.const 63
      i32.and
      i32.const 128
      i32.or
      i32.store8 offset=15
      local.get 2
      local.get 1
      i32.const 18
      i32.shr_u
      i32.const 240
      i32.or
      i32.store8 offset=12
      local.get 2
      local.get 1
      i32.const 6
      i32.shr_u
      i32.const 63
      i32.and
      i32.const 128
      i32.or
      i32.store8 offset=14
      local.get 2
      local.get 1
      i32.const 12
      i32.shr_u
      i32.const 63
      i32.and
      i32.const 128
      i32.or
      i32.store8 offset=13
      i32.const 4
      local.set 1
    end
    local.get 0
    local.get 2
    i32.const 12
    i32.add
    local.get 1
    call $_ZN68_$LT$core..fmt..builders..PadAdapter$u20$as$u20$core..fmt..Write$GT$9write_str17h7d69c319e62567eaE
    local.set 1
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 1)
  (func $_ZN4core3fmt5Write9write_fmt17hd43133c0fc77f05dE (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.store offset=4
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 1
    i64.load align=4
    i64.store offset=8
    local.get 2
    i32.const 4
    i32.add
    i32.const 1051096
    local.get 2
    i32.const 8
    i32.add
    call $_ZN4core3fmt5write17h8cfd01c67a4a46c9E
    local.set 1
    local.get 2
    i32.const 32
    i32.add
    global.set 0
    local.get 1)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17h0573ddc37a895899E (type 3) (param i32 i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 1
    local.get 2
    call $_ZN68_$LT$core..fmt..builders..PadAdapter$u20$as$u20$core..fmt..Write$GT$9write_str17h7d69c319e62567eaE)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17hc454648cbae23e1cE (type 4) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 1
    call $_ZN4core3fmt5Write10write_char17he6255f3eabfceb48E)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17hc0aee806cb052974E (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.load
    i32.store offset=4
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 1
    i64.load align=4
    i64.store offset=8
    local.get 2
    i32.const 4
    i32.add
    i32.const 1051096
    local.get 2
    i32.const 8
    i32.add
    call $_ZN4core3fmt5write17h8cfd01c67a4a46c9E
    local.set 1
    local.get 2
    i32.const 32
    i32.add
    global.set 0
    local.get 1)
  (func $_ZN4core3fmt10ArgumentV110show_usize17hf8ebfbd77fa8a93dE (type 4) (param i32 i32) (result i32)
    local.get 0
    i64.load32_u
    i32.const 1
    local.get 1
    call $_ZN4core3fmt3num3imp7fmt_u6417h4648b2300e699242E)
  (func $_ZN4core3fmt3num3imp7fmt_u6417h4648b2300e699242E (type 13) (param i64 i32 i32) (result i32)
    (local i32 i32 i32 i64 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 3
    global.set 0
    i32.const 39
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i64.const 10000
        i64.lt_u
        br_if 0 (;@2;)
        i32.const 39
        local.set 4
        loop  ;; label = @3
          local.get 3
          i32.const 9
          i32.add
          local.get 4
          i32.add
          local.tee 5
          i32.const -4
          i32.add
          local.get 0
          local.get 0
          i64.const 10000
          i64.div_u
          local.tee 6
          i64.const 10000
          i64.mul
          i64.sub
          i32.wrap_i64
          local.tee 7
          i32.const 100
          i32.div_u
          local.tee 8
          i32.const 1
          i32.shl
          i32.const 1050854
          i32.add
          i32.load16_u align=1
          i32.store16 align=1
          local.get 5
          i32.const -2
          i32.add
          local.get 7
          local.get 8
          i32.const 100
          i32.mul
          i32.sub
          i32.const 1
          i32.shl
          i32.const 1050854
          i32.add
          i32.load16_u align=1
          i32.store16 align=1
          local.get 4
          i32.const -4
          i32.add
          local.set 4
          local.get 0
          i64.const 99999999
          i64.gt_u
          local.set 5
          local.get 6
          local.set 0
          local.get 5
          br_if 0 (;@3;)
          br 2 (;@1;)
        end
      end
      local.get 0
      local.set 6
    end
    block  ;; label = @1
      local.get 6
      i32.wrap_i64
      local.tee 5
      i32.const 99
      i32.le_s
      br_if 0 (;@1;)
      local.get 3
      i32.const 9
      i32.add
      local.get 4
      i32.const -2
      i32.add
      local.tee 4
      i32.add
      local.get 6
      i32.wrap_i64
      local.tee 5
      local.get 5
      i32.const 65535
      i32.and
      i32.const 100
      i32.div_u
      local.tee 5
      i32.const 100
      i32.mul
      i32.sub
      i32.const 65535
      i32.and
      i32.const 1
      i32.shl
      i32.const 1050854
      i32.add
      i32.load16_u align=1
      i32.store16 align=1
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 5
        i32.const 9
        i32.gt_s
        br_if 0 (;@2;)
        local.get 3
        i32.const 9
        i32.add
        local.get 4
        i32.const -1
        i32.add
        local.tee 4
        i32.add
        local.get 5
        i32.const 48
        i32.add
        i32.store8
        br 1 (;@1;)
      end
      local.get 3
      i32.const 9
      i32.add
      local.get 4
      i32.const -2
      i32.add
      local.tee 4
      i32.add
      local.get 5
      i32.const 1
      i32.shl
      i32.const 1050854
      i32.add
      i32.load16_u align=1
      i32.store16 align=1
    end
    local.get 2
    local.get 1
    i32.const 1050184
    i32.const 0
    local.get 3
    i32.const 9
    i32.add
    local.get 4
    i32.add
    i32.const 39
    local.get 4
    i32.sub
    call $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE
    local.set 4
    local.get 3
    i32.const 48
    i32.add
    global.set 0
    local.get 4)
  (func $_ZN59_$LT$core..fmt..Arguments$u20$as$u20$core..fmt..Display$GT$3fmt17he3a12226fd5f8b44E (type 4) (param i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 1
    i32.const 28
    i32.add
    i32.load
    local.set 3
    local.get 1
    i32.load offset=24
    local.set 1
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    local.get 0
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    local.get 0
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 0
    i64.load align=4
    i64.store offset=8
    local.get 1
    local.get 3
    local.get 2
    i32.const 8
    i32.add
    call $_ZN4core3fmt5write17h8cfd01c67a4a46c9E
    local.set 0
    local.get 2
    i32.const 32
    i32.add
    global.set 0
    local.get 0)
  (func $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE (type 14) (param i32 i32 i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.eqz
        br_if 0 (;@2;)
        i32.const 43
        i32.const 1114112
        local.get 0
        i32.load
        local.tee 6
        i32.const 1
        i32.and
        local.tee 1
        select
        local.set 7
        local.get 1
        local.get 5
        i32.add
        local.set 8
        br 1 (;@1;)
      end
      local.get 5
      i32.const 1
      i32.add
      local.set 8
      local.get 0
      i32.load
      local.set 6
      i32.const 45
      local.set 7
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 4
        i32.and
        br_if 0 (;@2;)
        i32.const 0
        local.set 2
        br 1 (;@1;)
      end
      i32.const 0
      local.set 9
      block  ;; label = @2
        local.get 3
        i32.eqz
        br_if 0 (;@2;)
        local.get 3
        local.set 10
        local.get 2
        local.set 1
        loop  ;; label = @3
          local.get 9
          local.get 1
          i32.load8_u
          i32.const 192
          i32.and
          i32.const 128
          i32.eq
          i32.add
          local.set 9
          local.get 1
          i32.const 1
          i32.add
          local.set 1
          local.get 10
          i32.const -1
          i32.add
          local.tee 10
          br_if 0 (;@3;)
        end
      end
      local.get 8
      local.get 3
      i32.add
      local.get 9
      i32.sub
      local.set 8
    end
    i32.const 1
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                local.get 0
                                i32.load offset=8
                                i32.const 1
                                i32.ne
                                br_if 0 (;@14;)
                                local.get 0
                                i32.const 12
                                i32.add
                                i32.load
                                local.tee 9
                                local.get 8
                                i32.le_u
                                br_if 1 (;@13;)
                                local.get 6
                                i32.const 8
                                i32.and
                                br_if 2 (;@12;)
                                local.get 9
                                local.get 8
                                i32.sub
                                local.set 1
                                i32.const 1
                                local.get 0
                                i32.load8_u offset=48
                                local.tee 9
                                local.get 9
                                i32.const 3
                                i32.eq
                                select
                                local.tee 9
                                i32.const 3
                                i32.and
                                i32.eqz
                                br_if 3 (;@11;)
                                local.get 9
                                i32.const 2
                                i32.eq
                                br_if 4 (;@10;)
                                i32.const 0
                                local.set 11
                                local.get 1
                                local.set 9
                                br 5 (;@9;)
                              end
                              local.get 0
                              local.get 7
                              local.get 2
                              local.get 3
                              call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17h80e5cb0c85711a72E
                              br_if 12 (;@1;)
                              local.get 0
                              i32.load offset=24
                              local.get 4
                              local.get 5
                              local.get 0
                              i32.const 28
                              i32.add
                              i32.load
                              i32.load offset=12
                              call_indirect (type 3)
                              return
                            end
                            local.get 0
                            local.get 7
                            local.get 2
                            local.get 3
                            call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17h80e5cb0c85711a72E
                            br_if 11 (;@1;)
                            local.get 0
                            i32.load offset=24
                            local.get 4
                            local.get 5
                            local.get 0
                            i32.const 28
                            i32.add
                            i32.load
                            i32.load offset=12
                            call_indirect (type 3)
                            return
                          end
                          i32.const 1
                          local.set 1
                          local.get 0
                          i32.const 1
                          i32.store8 offset=48
                          local.get 0
                          i32.const 48
                          i32.store offset=4
                          local.get 0
                          local.get 7
                          local.get 2
                          local.get 3
                          call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17h80e5cb0c85711a72E
                          br_if 10 (;@1;)
                          local.get 9
                          local.get 8
                          i32.sub
                          local.set 1
                          i32.const 1
                          local.get 0
                          i32.const 48
                          i32.add
                          i32.load8_u
                          local.tee 9
                          local.get 9
                          i32.const 3
                          i32.eq
                          select
                          local.tee 9
                          i32.const 3
                          i32.and
                          i32.eqz
                          br_if 3 (;@8;)
                          local.get 9
                          i32.const 2
                          i32.eq
                          br_if 4 (;@7;)
                          i32.const 0
                          local.set 8
                          local.get 1
                          local.set 9
                          br 5 (;@6;)
                        end
                        i32.const 0
                        local.set 9
                        local.get 1
                        local.set 11
                        br 1 (;@9;)
                      end
                      local.get 1
                      i32.const 1
                      i32.shr_u
                      local.set 9
                      local.get 1
                      i32.const 1
                      i32.add
                      i32.const 1
                      i32.shr_u
                      local.set 11
                    end
                    i32.const -1
                    local.set 1
                    local.get 0
                    i32.const 4
                    i32.add
                    local.set 10
                    local.get 0
                    i32.const 24
                    i32.add
                    local.set 8
                    local.get 0
                    i32.const 28
                    i32.add
                    local.set 6
                    block  ;; label = @9
                      loop  ;; label = @10
                        local.get 1
                        i32.const 1
                        i32.add
                        local.tee 1
                        local.get 9
                        i32.ge_u
                        br_if 1 (;@9;)
                        local.get 8
                        i32.load
                        local.get 10
                        i32.load
                        local.get 6
                        i32.load
                        i32.load offset=16
                        call_indirect (type 4)
                        i32.eqz
                        br_if 0 (;@10;)
                        br 5 (;@5;)
                      end
                    end
                    local.get 0
                    i32.const 4
                    i32.add
                    i32.load
                    local.set 10
                    i32.const 1
                    local.set 1
                    local.get 0
                    local.get 7
                    local.get 2
                    local.get 3
                    call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17h80e5cb0c85711a72E
                    br_if 7 (;@1;)
                    local.get 0
                    i32.const 24
                    i32.add
                    local.tee 9
                    i32.load
                    local.get 4
                    local.get 5
                    local.get 0
                    i32.const 28
                    i32.add
                    local.tee 3
                    i32.load
                    i32.load offset=12
                    call_indirect (type 3)
                    br_if 7 (;@1;)
                    local.get 9
                    i32.load
                    local.set 0
                    i32.const -1
                    local.set 9
                    local.get 3
                    i32.load
                    i32.const 16
                    i32.add
                    local.set 3
                    loop  ;; label = @9
                      local.get 9
                      i32.const 1
                      i32.add
                      local.tee 9
                      local.get 11
                      i32.ge_u
                      br_if 6 (;@3;)
                      i32.const 1
                      local.set 1
                      local.get 0
                      local.get 10
                      local.get 3
                      i32.load
                      call_indirect (type 4)
                      i32.eqz
                      br_if 0 (;@9;)
                      br 8 (;@1;)
                    end
                  end
                  i32.const 0
                  local.set 9
                  local.get 1
                  local.set 8
                  br 1 (;@6;)
                end
                local.get 1
                i32.const 1
                i32.shr_u
                local.set 9
                local.get 1
                i32.const 1
                i32.add
                i32.const 1
                i32.shr_u
                local.set 8
              end
              i32.const -1
              local.set 1
              local.get 0
              i32.const 4
              i32.add
              local.set 10
              local.get 0
              i32.const 24
              i32.add
              local.set 3
              local.get 0
              i32.const 28
              i32.add
              local.set 2
              loop  ;; label = @6
                local.get 1
                i32.const 1
                i32.add
                local.tee 1
                local.get 9
                i32.ge_u
                br_if 2 (;@4;)
                local.get 3
                i32.load
                local.get 10
                i32.load
                local.get 2
                i32.load
                i32.load offset=16
                call_indirect (type 4)
                i32.eqz
                br_if 0 (;@6;)
              end
            end
            i32.const 1
            local.set 1
            br 3 (;@1;)
          end
          local.get 0
          i32.const 4
          i32.add
          i32.load
          local.set 10
          i32.const 1
          local.set 1
          local.get 0
          i32.const 24
          i32.add
          local.tee 9
          i32.load
          local.get 4
          local.get 5
          local.get 0
          i32.const 28
          i32.add
          local.tee 3
          i32.load
          i32.load offset=12
          call_indirect (type 3)
          br_if 2 (;@1;)
          local.get 9
          i32.load
          local.set 0
          i32.const -1
          local.set 9
          local.get 3
          i32.load
          i32.const 16
          i32.add
          local.set 3
          loop  ;; label = @4
            local.get 9
            i32.const 1
            i32.add
            local.tee 9
            local.get 8
            i32.ge_u
            br_if 2 (;@2;)
            i32.const 1
            local.set 1
            local.get 0
            local.get 10
            local.get 3
            i32.load
            call_indirect (type 4)
            i32.eqz
            br_if 0 (;@4;)
            br 3 (;@1;)
          end
        end
        i32.const 0
        return
      end
      i32.const 0
      return
    end
    local.get 1)
  (func $_ZN4core3fmt9Formatter12pad_integral12write_prefix17h80e5cb0c85711a72E (type 10) (param i32 i32 i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.const 1114112
          i32.eq
          br_if 0 (;@3;)
          i32.const 1
          local.set 4
          local.get 0
          i32.load offset=24
          local.get 1
          local.get 0
          i32.const 28
          i32.add
          i32.load
          i32.load offset=16
          call_indirect (type 4)
          br_if 1 (;@2;)
        end
        local.get 2
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        i32.load offset=24
        local.get 2
        local.get 3
        local.get 0
        i32.const 28
        i32.add
        i32.load
        i32.load offset=12
        call_indirect (type 3)
        local.set 4
      end
      local.get 4
      return
    end
    i32.const 0)
  (func $_ZN4core3fmt9Formatter15debug_lower_hex17hfff82ceb2b4262ffE (type 2) (param i32) (result i32)
    local.get 0
    i32.load8_u
    i32.const 16
    i32.and
    i32.const 4
    i32.shr_u)
  (func $_ZN4core3fmt9Formatter15debug_upper_hex17h0a841ca4788fd580E (type 2) (param i32) (result i32)
    local.get 0
    i32.load8_u
    i32.const 32
    i32.and
    i32.const 5
    i32.shr_u)
  (func $_ZN4core3fmt9Formatter11debug_tuple17h3d2072b8cea76c89E (type 9) (param i32 i32 i32 i32)
    local.get 0
    local.get 1
    i32.load offset=24
    local.get 2
    local.get 3
    local.get 1
    i32.const 28
    i32.add
    i32.load
    i32.load offset=12
    call_indirect (type 3)
    i32.store8 offset=8
    local.get 0
    local.get 1
    i32.store
    local.get 0
    local.get 3
    i32.eqz
    i32.store8 offset=9
    local.get 0
    i32.const 0
    i32.store offset=4)
  (func $_ZN4core3fmt9Formatter10debug_list17h2471317d64372e8dE (type 1) (param i32 i32)
    (local i32)
    local.get 1
    i32.load offset=24
    i32.const 1051094
    i32.const 1
    local.get 1
    i32.const 28
    i32.add
    i32.load
    i32.load offset=12
    call_indirect (type 3)
    local.set 2
    local.get 0
    i32.const 0
    i32.store8 offset=5
    local.get 0
    local.get 2
    i32.store8 offset=4
    local.get 0
    local.get 1
    i32.store)
  (func $_ZN40_$LT$str$u20$as$u20$core..fmt..Debug$GT$3fmt17h5b1f8bd45e1c8428E (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 3
    global.set 0
    i32.const 1
    local.set 4
    block  ;; label = @1
      local.get 2
      i32.load offset=24
      i32.const 34
      local.get 2
      i32.const 28
      i32.add
      i32.load
      i32.load offset=16
      call_indirect (type 4)
      br_if 0 (;@1;)
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          local.get 1
          i32.add
          local.set 5
          local.get 2
          i32.const 24
          i32.add
          local.set 6
          local.get 2
          i32.const 28
          i32.add
          local.set 7
          local.get 0
          local.set 8
          i32.const 0
          local.set 9
          i32.const 0
          local.set 10
          block  ;; label = @4
            loop  ;; label = @5
              local.get 8
              local.set 11
              local.get 8
              i32.const 1
              i32.add
              local.set 12
              block  ;; label = @6
                block  ;; label = @7
                  local.get 8
                  i32.load8_s
                  local.tee 13
                  i32.const 0
                  i32.lt_s
                  br_if 0 (;@7;)
                  local.get 13
                  i32.const 255
                  i32.and
                  local.set 13
                  local.get 12
                  local.set 8
                  br 1 (;@6;)
                end
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 12
                    local.get 5
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 12
                    i32.load8_u
                    i32.const 63
                    i32.and
                    local.set 14
                    local.get 8
                    i32.const 2
                    i32.add
                    local.tee 8
                    local.set 12
                    br 1 (;@7;)
                  end
                  i32.const 0
                  local.set 14
                  local.get 5
                  local.set 8
                end
                local.get 13
                i32.const 31
                i32.and
                local.set 15
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 13
                      i32.const 255
                      i32.and
                      local.tee 13
                      i32.const 224
                      i32.lt_u
                      br_if 0 (;@9;)
                      local.get 8
                      local.get 5
                      i32.eq
                      br_if 1 (;@8;)
                      local.get 8
                      i32.load8_u
                      i32.const 63
                      i32.and
                      local.set 16
                      local.get 8
                      i32.const 1
                      i32.add
                      local.tee 12
                      local.set 17
                      br 2 (;@7;)
                    end
                    local.get 14
                    local.get 15
                    i32.const 6
                    i32.shl
                    i32.or
                    local.set 13
                    local.get 12
                    local.set 8
                    br 2 (;@6;)
                  end
                  i32.const 0
                  local.set 16
                  local.get 5
                  local.set 17
                end
                local.get 16
                local.get 14
                i32.const 6
                i32.shl
                i32.or
                local.set 14
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 13
                      i32.const 240
                      i32.lt_u
                      br_if 0 (;@9;)
                      local.get 17
                      local.get 5
                      i32.eq
                      br_if 1 (;@8;)
                      local.get 17
                      i32.const 1
                      i32.add
                      local.set 8
                      local.get 17
                      i32.load8_u
                      i32.const 63
                      i32.and
                      local.set 13
                      br 2 (;@7;)
                    end
                    local.get 14
                    local.get 15
                    i32.const 12
                    i32.shl
                    i32.or
                    local.set 13
                    local.get 12
                    local.set 8
                    br 2 (;@6;)
                  end
                  i32.const 0
                  local.set 13
                  local.get 12
                  local.set 8
                end
                local.get 14
                i32.const 6
                i32.shl
                local.get 15
                i32.const 18
                i32.shl
                i32.const 1835008
                i32.and
                i32.or
                local.get 13
                i32.or
                local.tee 13
                i32.const 1114112
                i32.eq
                br_if 2 (;@4;)
              end
              i32.const 2
              local.set 12
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                local.get 13
                                i32.const -9
                                i32.add
                                local.tee 15
                                i32.const 30
                                i32.gt_u
                                br_if 0 (;@14;)
                                i32.const 116
                                local.set 14
                                block  ;; label = @15
                                  local.get 15
                                  br_table 8 (;@7;) 0 (;@15;) 3 (;@12;) 3 (;@12;) 4 (;@11;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 2 (;@13;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 3 (;@12;) 2 (;@13;) 8 (;@7;)
                                end
                                i32.const 110
                                local.set 14
                                br 4 (;@10;)
                              end
                              local.get 13
                              i32.const 92
                              i32.ne
                              br_if 1 (;@12;)
                            end
                            local.get 13
                            local.set 14
                            br 5 (;@7;)
                          end
                          i32.const 1054144
                          local.get 13
                          call $_ZN4core7unicode9bool_trie8BoolTrie6lookup17h03e59958127816b1E
                          i32.eqz
                          br_if 2 (;@9;)
                          br 3 (;@8;)
                        end
                        i32.const 114
                        local.set 14
                      end
                      br 2 (;@7;)
                    end
                    block  ;; label = @9
                      local.get 13
                      i32.const 65535
                      i32.gt_u
                      br_if 0 (;@9;)
                      local.get 13
                      i32.const 1051296
                      i32.const 40
                      i32.const 1051376
                      i32.const 303
                      i32.const 1051679
                      i32.const 316
                      call $_ZN4core7unicode9printable5check17hc99709d493feb50bE
                      i32.eqz
                      br_if 1 (;@8;)
                      br 3 (;@6;)
                    end
                    block  ;; label = @9
                      local.get 13
                      i32.const 131071
                      i32.gt_u
                      br_if 0 (;@9;)
                      local.get 13
                      i32.const 1051995
                      i32.const 33
                      i32.const 1052061
                      i32.const 158
                      i32.const 1052219
                      i32.const 381
                      call $_ZN4core7unicode9printable5check17hc99709d493feb50bE
                      i32.eqz
                      br_if 1 (;@8;)
                      br 3 (;@6;)
                    end
                    local.get 13
                    i32.const 917999
                    i32.gt_u
                    br_if 0 (;@8;)
                    local.get 13
                    i32.const -195102
                    i32.add
                    i32.const 722658
                    i32.lt_u
                    br_if 0 (;@8;)
                    local.get 13
                    i32.const -191457
                    i32.add
                    i32.const 3103
                    i32.lt_u
                    br_if 0 (;@8;)
                    local.get 13
                    i32.const -183970
                    i32.add
                    i32.const 14
                    i32.lt_u
                    br_if 0 (;@8;)
                    local.get 13
                    i32.const 2097150
                    i32.and
                    i32.const 178206
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 13
                    i32.const -173783
                    i32.add
                    i32.const 41
                    i32.lt_u
                    br_if 0 (;@8;)
                    local.get 13
                    i32.const -177973
                    i32.add
                    i32.const 10
                    i32.gt_u
                    br_if 2 (;@6;)
                  end
                  local.get 13
                  i32.const 1
                  i32.or
                  i32.clz
                  i32.const 2
                  i32.shr_u
                  i32.const 7
                  i32.xor
                  i64.extend_i32_u
                  i64.const 21474836480
                  i64.or
                  local.set 18
                  i32.const 3
                  local.set 12
                  local.get 13
                  local.set 14
                end
                local.get 3
                local.get 1
                i32.store offset=4
                local.get 3
                local.get 0
                i32.store
                local.get 3
                local.get 9
                i32.store offset=8
                local.get 3
                local.get 10
                i32.store offset=12
                block  ;; label = @7
                  local.get 10
                  local.get 9
                  i32.lt_u
                  br_if 0 (;@7;)
                  block  ;; label = @8
                    local.get 9
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 9
                    local.get 1
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 9
                    local.get 1
                    i32.ge_u
                    br_if 1 (;@7;)
                    local.get 0
                    local.get 9
                    i32.add
                    i32.load8_s
                    i32.const -65
                    i32.le_s
                    br_if 1 (;@7;)
                  end
                  block  ;; label = @8
                    local.get 10
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 10
                    local.get 1
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 10
                    local.get 1
                    i32.ge_u
                    br_if 1 (;@7;)
                    local.get 0
                    local.get 10
                    i32.add
                    i32.load8_s
                    i32.const -65
                    i32.le_s
                    br_if 1 (;@7;)
                  end
                  block  ;; label = @8
                    local.get 6
                    i32.load
                    local.get 0
                    local.get 9
                    i32.add
                    local.get 10
                    local.get 9
                    i32.sub
                    local.get 7
                    i32.load
                    i32.load offset=12
                    call_indirect (type 3)
                    br_if 0 (;@8;)
                    loop  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    local.get 12
                                    i32.const 1
                                    i32.eq
                                    br_if 0 (;@16;)
                                    i32.const 92
                                    local.set 9
                                    block  ;; label = @17
                                      local.get 12
                                      i32.const 2
                                      i32.eq
                                      br_if 0 (;@17;)
                                      local.get 12
                                      i32.const 3
                                      i32.ne
                                      br_if 6 (;@11;)
                                      local.get 18
                                      i64.const 32
                                      i64.shr_u
                                      i32.wrap_i64
                                      i32.const 255
                                      i32.and
                                      i32.const -1
                                      i32.add
                                      local.tee 12
                                      i32.const 4
                                      i32.gt_u
                                      br_if 6 (;@11;)
                                      block  ;; label = @18
                                        local.get 12
                                        br_table 0 (;@18;) 6 (;@12;) 4 (;@14;) 5 (;@13;) 3 (;@15;) 0 (;@18;)
                                      end
                                      local.get 18
                                      i64.const -1095216660481
                                      i64.and
                                      local.set 18
                                      i32.const 3
                                      local.set 12
                                      i32.const 125
                                      local.set 9
                                      br 7 (;@10;)
                                    end
                                    i32.const 1
                                    local.set 12
                                    br 6 (;@10;)
                                  end
                                  i32.const 0
                                  local.set 12
                                  local.get 14
                                  local.set 9
                                  br 5 (;@10;)
                                end
                                local.get 18
                                i64.const -1095216660481
                                i64.and
                                i64.const 17179869184
                                i64.or
                                local.set 18
                                i32.const 3
                                local.set 12
                                br 4 (;@10;)
                              end
                              local.get 18
                              i64.const -1095216660481
                              i64.and
                              i64.const 8589934592
                              i64.or
                              local.set 18
                              i32.const 3
                              local.set 12
                              i32.const 123
                              local.set 9
                              br 3 (;@10;)
                            end
                            local.get 18
                            i64.const -1095216660481
                            i64.and
                            i64.const 12884901888
                            i64.or
                            local.set 18
                            i32.const 3
                            local.set 12
                            i32.const 117
                            local.set 9
                            br 2 (;@10;)
                          end
                          local.get 14
                          local.get 18
                          i32.wrap_i64
                          local.tee 15
                          i32.const 2
                          i32.shl
                          i32.const 28
                          i32.and
                          i32.shr_u
                          i32.const 15
                          i32.and
                          local.tee 12
                          i32.const 48
                          i32.or
                          local.get 12
                          i32.const 87
                          i32.add
                          local.get 12
                          i32.const 10
                          i32.lt_u
                          select
                          local.set 9
                          block  ;; label = @12
                            local.get 15
                            i32.eqz
                            br_if 0 (;@12;)
                            local.get 18
                            i64.const -1
                            i64.add
                            i64.const 4294967295
                            i64.and
                            local.get 18
                            i64.const -4294967296
                            i64.and
                            i64.or
                            local.set 18
                            i32.const 3
                            local.set 12
                            br 2 (;@10;)
                          end
                          local.get 18
                          i64.const -1095216660481
                          i64.and
                          i64.const 4294967296
                          i64.or
                          local.set 18
                          i32.const 3
                          local.set 12
                          br 1 (;@10;)
                        end
                        i32.const 1
                        local.set 12
                        block  ;; label = @11
                          local.get 13
                          i32.const 128
                          i32.lt_u
                          br_if 0 (;@11;)
                          i32.const 2
                          local.set 12
                          local.get 13
                          i32.const 2048
                          i32.lt_u
                          br_if 0 (;@11;)
                          i32.const 3
                          i32.const 4
                          local.get 13
                          i32.const 65536
                          i32.lt_u
                          select
                          local.set 12
                        end
                        local.get 12
                        local.get 10
                        i32.add
                        local.set 9
                        br 4 (;@6;)
                      end
                      local.get 6
                      i32.load
                      local.get 9
                      local.get 7
                      i32.load
                      i32.load offset=16
                      call_indirect (type 4)
                      i32.eqz
                      br_if 0 (;@9;)
                    end
                  end
                  i32.const 1
                  local.set 4
                  br 6 (;@1;)
                end
                local.get 3
                local.get 3
                i32.const 12
                i32.add
                i32.store offset=24
                local.get 3
                local.get 3
                i32.const 8
                i32.add
                i32.store offset=20
                local.get 3
                local.get 3
                i32.store offset=16
                local.get 3
                i32.const 16
                i32.add
                call $_ZN4core3str6traits101_$LT$impl$u20$core..slice..SliceIndex$LT$str$GT$$u20$for$u20$core..ops..range..Range$LT$usize$GT$$GT$5index28_$u7b$$u7b$closure$u7d$$u7d$17h5fc12def9a7684b5E
                unreachable
              end
              local.get 10
              local.get 11
              i32.sub
              local.get 8
              i32.add
              local.set 10
              local.get 5
              local.get 8
              i32.ne
              br_if 0 (;@5;)
            end
          end
          local.get 9
          i32.eqz
          br_if 1 (;@2;)
          local.get 9
          local.get 1
          i32.eq
          br_if 1 (;@2;)
          block  ;; label = @4
            local.get 9
            local.get 1
            i32.ge_u
            br_if 0 (;@4;)
            local.get 0
            local.get 9
            i32.add
            i32.load8_s
            i32.const -65
            i32.gt_s
            br_if 2 (;@2;)
          end
          local.get 0
          local.get 1
          local.get 9
          local.get 1
          call $_ZN4core3str16slice_error_fail17h3704ce74b976be71E
          unreachable
        end
        i32.const 0
        local.set 9
      end
      local.get 2
      i32.const 24
      i32.add
      local.tee 12
      i32.load
      local.get 0
      local.get 9
      i32.add
      local.get 1
      local.get 9
      i32.sub
      local.get 2
      i32.const 28
      i32.add
      local.tee 10
      i32.load
      i32.load offset=12
      call_indirect (type 3)
      br_if 0 (;@1;)
      local.get 12
      i32.load
      i32.const 34
      local.get 10
      i32.load
      i32.load offset=16
      call_indirect (type 4)
      local.set 4
    end
    local.get 3
    i32.const 32
    i32.add
    global.set 0
    local.get 4)
  (func $_ZN42_$LT$str$u20$as$u20$core..fmt..Display$GT$3fmt17hfdf78a4ae71e6a9bE (type 3) (param i32 i32 i32) (result i32)
    local.get 2
    local.get 0
    local.get 1
    call $_ZN4core3fmt9Formatter3pad17hacaebd5abd28adf1E)
  (func $_ZN4core3fmt3num53_$LT$impl$u20$core..fmt..LowerHex$u20$for$u20$i32$GT$3fmt17h92e35f33f4a021aeE (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 128
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    i32.load
    local.set 3
    i32.const 0
    local.set 0
    loop  ;; label = @1
      local.get 2
      local.get 0
      i32.add
      i32.const 127
      i32.add
      local.get 3
      i32.const 15
      i32.and
      local.tee 4
      i32.const 48
      i32.or
      local.get 4
      i32.const 87
      i32.add
      local.get 4
      i32.const 10
      i32.lt_u
      select
      i32.store8
      local.get 0
      i32.const -1
      i32.add
      local.set 0
      local.get 3
      i32.const 4
      i32.shr_u
      local.tee 3
      br_if 0 (;@1;)
    end
    block  ;; label = @1
      local.get 0
      i32.const 128
      i32.add
      local.tee 3
      i32.const 129
      i32.ge_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 1
      i32.const 1050852
      i32.const 2
      local.get 2
      local.get 0
      i32.add
      i32.const 128
      i32.add
      i32.const 0
      local.get 0
      i32.sub
      call $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE
      local.set 0
      local.get 2
      i32.const 128
      i32.add
      global.set 0
      local.get 0
      return
    end
    local.get 3
    i32.const 128
    call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
    unreachable)
  (func $_ZN4core3fmt3num52_$LT$impl$u20$core..fmt..UpperHex$u20$for$u20$i8$GT$3fmt17hb957d38375b60145E (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 128
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    i32.load8_u
    local.set 3
    i32.const 0
    local.set 0
    loop  ;; label = @1
      local.get 2
      local.get 0
      i32.add
      i32.const 127
      i32.add
      local.get 3
      i32.const 15
      i32.and
      local.tee 4
      i32.const 48
      i32.or
      local.get 4
      i32.const 55
      i32.add
      local.get 4
      i32.const 10
      i32.lt_u
      select
      i32.store8
      local.get 0
      i32.const -1
      i32.add
      local.set 0
      local.get 3
      i32.const 4
      i32.shr_u
      i32.const 15
      i32.and
      local.tee 3
      br_if 0 (;@1;)
    end
    block  ;; label = @1
      local.get 0
      i32.const 128
      i32.add
      local.tee 3
      i32.const 129
      i32.ge_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 1
      i32.const 1050852
      i32.const 2
      local.get 2
      local.get 0
      i32.add
      i32.const 128
      i32.add
      i32.const 0
      local.get 0
      i32.sub
      call $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE
      local.set 0
      local.get 2
      i32.const 128
      i32.add
      global.set 0
      local.get 0
      return
    end
    local.get 3
    i32.const 128
    call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
    unreachable)
  (func $_ZN4core3fmt3num53_$LT$impl$u20$core..fmt..UpperHex$u20$for$u20$i32$GT$3fmt17he210bef9ff0aa51fE (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 128
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    i32.load
    local.set 3
    i32.const 0
    local.set 0
    loop  ;; label = @1
      local.get 2
      local.get 0
      i32.add
      i32.const 127
      i32.add
      local.get 3
      i32.const 15
      i32.and
      local.tee 4
      i32.const 48
      i32.or
      local.get 4
      i32.const 55
      i32.add
      local.get 4
      i32.const 10
      i32.lt_u
      select
      i32.store8
      local.get 0
      i32.const -1
      i32.add
      local.set 0
      local.get 3
      i32.const 4
      i32.shr_u
      local.tee 3
      br_if 0 (;@1;)
    end
    block  ;; label = @1
      local.get 0
      i32.const 128
      i32.add
      local.tee 3
      i32.const 129
      i32.ge_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 1
      i32.const 1050852
      i32.const 2
      local.get 2
      local.get 0
      i32.add
      i32.const 128
      i32.add
      i32.const 0
      local.get 0
      i32.sub
      call $_ZN4core3fmt9Formatter12pad_integral17h17dddcbb38d9710fE
      local.set 0
      local.get 2
      i32.const 128
      i32.add
      global.set 0
      local.get 0
      return
    end
    local.get 3
    i32.const 128
    call $_ZN4core5slice22slice_index_order_fail17h719e3c8ae46c9d99E
    unreachable)
  (func $memcpy (type 3) (param i32 i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.set 3
      loop  ;; label = @2
        local.get 3
        local.get 1
        i32.load8_u
        i32.store8
        local.get 3
        i32.const 1
        i32.add
        local.set 3
        local.get 1
        i32.const 1
        i32.add
        local.set 1
        local.get 2
        i32.const -1
        i32.add
        local.tee 2
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func $memcmp (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        local.set 3
        loop  ;; label = @3
          local.get 0
          local.get 3
          i32.add
          i32.load8_u
          local.tee 4
          local.get 1
          local.get 3
          i32.add
          i32.load8_u
          local.tee 5
          i32.ne
          br_if 2 (;@1;)
          local.get 3
          i32.const 1
          i32.add
          local.tee 3
          local.get 2
          i32.lt_u
          br_if 0 (;@3;)
        end
        i32.const 0
        return
      end
      i32.const 0
      return
    end
    local.get 4
    local.get 5
    i32.sub)
  (func $inc_and_get (type 2) (param i32) (result i32)
    local.get 0
    call $inc
    call $get)
  (func $copy_from_reg_and_get (type 2) (param i32) (result i32)
    local.get 0
    call $copy_from_reg
    call $get)
  (func $copy_to_reg_and_get (type 2) (param i32) (result i32)
    local.get 0
    call $copy_to_reg
    call $get)
  (func $_ZN7counter4main17ha38116a1f7c42429E (type 5))
  (func $main (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 45
    i32.store offset=12
    local.get 2
    i32.const 12
    i32.add
    i32.const 1055672
    local.get 0
    local.get 1
    call $_ZN3std2rt19lang_start_internal17he5cddafed9dd65a0E
    local.set 0
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func $_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h8505eb24e9b3d0b1E (type 2) (param i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    i32.load
    call_indirect (type 5)
    local.get 1
    i32.const 0
    i32.store8 offset=15
    local.get 1
    i32.const 15
    i32.add
    call $_ZN3std3sys4wasm7process8ExitCode6as_i3217h8c3f6c9d0d4bbc02E
    local.set 0
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func $_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17he4866c7624714632E (type 2) (param i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    i32.load
    call_indirect (type 5)
    local.get 1
    i32.const 0
    i32.store8 offset=15
    local.get 1
    i32.const 15
    i32.add
    call $_ZN3std3sys4wasm7process8ExitCode6as_i3217h8c3f6c9d0d4bbc02E
    local.set 0
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func $_ZN4core3ptr18real_drop_in_place17hd95bbad20902bbffE (type 0) (param i32))
  (func $__rust_alloc (type 4) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call $__rdl_alloc)
  (func $__rust_dealloc (type 8) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call $__rdl_dealloc)
  (func $__rust_realloc (type 10) (param i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    call $__rdl_realloc)
  (table (;0;) 49 49 funcref)
  (memory (;0;) 17)
  (global (;0;) (mut i32) (i32.const 1048576))
  (global (;1;) i32 (i32.const 1056202))
  (global (;2;) i32 (i32.const 1056202))
  (export "memory" (memory 0))
  (export "__heap_base" (global 1))
  (export "__data_end" (global 2))
  (export "inc_and_get" (func $inc_and_get))
  (export "copy_from_reg_and_get" (func $copy_from_reg_and_get))
  (export "copy_to_reg_and_get" (func $copy_to_reg_and_get))
  (export "main" (func $main))
  (elem (;0;) (i32.const 1) $_ZN63_$LT$core..cell..BorrowMutError$u20$as$u20$core..fmt..Debug$GT$3fmt17h94d425ffcd3a62fcE $_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17haf2d0b9687871f57E $_ZN82_$LT$std..sys_common..poison..PoisonError$LT$T$GT$$u20$as$u20$core..fmt..Debug$GT$3fmt17h28d7bea72dfee2beE $_ZN62_$LT$std..ffi..c_str..NulError$u20$as$u20$core..fmt..Debug$GT$3fmt17h78d171950c7b08abE $_ZN60_$LT$core..cell..BorrowError$u20$as$u20$core..fmt..Debug$GT$3fmt17h1bceae61083d85a2E $_ZN59_$LT$core..fmt..Arguments$u20$as$u20$core..fmt..Display$GT$3fmt17he3a12226fd5f8b44E $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17hbcb23e4f920baed2E $_ZN3std5alloc24default_alloc_error_hook17h4dee2a5fdd1dba75E $_ZN3std9panicking3try7do_call17hb9a857ad47fc7875E $_ZN76_$LT$std..sys_common..thread_local..Key$u20$as$u20$core..ops..drop..Drop$GT$4drop17he3d36d5f1169eee9E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17h2ea757743c87a7b0E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17h6932f8aa37f23f2cE $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17hfbe1b10105aef519E $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h37cfe269a54a2708E $_ZN4core3ptr18real_drop_in_place17hc5a29f899ad0c715E $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h64cc3526f0a264d5E $_ZN3std4sync4once4Once9call_once28_$u7b$$u7b$closure$u7d$$u7d$17he0ac5916c2b898e8E $_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h5336016dc9200794E $_ZN4core3ptr18real_drop_in_place17h7615a9d826f89d4cE $_ZN89_$LT$std..panicking..continue_panic_fmt..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$9box_me_up17hf287a40b33981434E $_ZN89_$LT$std..panicking..continue_panic_fmt..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$3get17h0c4a906bac5cdb67E $_ZN4core3ptr18real_drop_in_place17h41544376cef8a15bE $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h4139b7e5fe9883cbE $_ZN4core3ptr18real_drop_in_place17h0372efb5da71b90dE $_ZN91_$LT$std..panicking..begin_panic..PanicPayload$LT$A$GT$$u20$as$u20$core..panic..BoxMeUp$GT$9box_me_up17h76250b59498caedbE $_ZN91_$LT$std..panicking..begin_panic..PanicPayload$LT$A$GT$$u20$as$u20$core..panic..BoxMeUp$GT$3get17h634a6ff81d071bd5E $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17he5dd53c4163a6f10E $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h27b62dc0d035ddf3E $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h3575d213ce7d211eE $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$u32$GT$3fmt17h4f3293445fab7cb7E $_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17hd813a4fb6bc448f6E $_ZN71_$LT$core..ops..range..Range$LT$Idx$GT$$u20$as$u20$core..fmt..Debug$GT$3fmt17h6105e754ea0aeafbE $_ZN41_$LT$char$u20$as$u20$core..fmt..Debug$GT$3fmt17h7a420490f85a5bf3E $_ZN4core3fmt10ArgumentV110show_usize17hf8ebfbd77fa8a93dE $_ZN4core3ptr18real_drop_in_place17h5a9a7fbf25605767E $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h5fb1d47f0acdabccE $_ZN4core3ptr18real_drop_in_place17h23040579a46f15a6E $_ZN68_$LT$core..fmt..builders..PadAdapter$u20$as$u20$core..fmt..Write$GT$9write_str17h7d69c319e62567eaE $_ZN4core3fmt5Write10write_char17he6255f3eabfceb48E $_ZN4core3fmt5Write9write_fmt17hd43133c0fc77f05dE $_ZN4core3ptr18real_drop_in_place17h0013f08692be52e6E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17h0573ddc37a895899E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17hc454648cbae23e1cE $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17hc0aee806cb052974E $_ZN7counter4main17ha38116a1f7c42429E $_ZN4core3ptr18real_drop_in_place17hd95bbad20902bbffE $_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h8505eb24e9b3d0b1E $_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17he4866c7624714632E)
  (data (;0;) (i32.const 1048576) "\0a\00\00\00\04\00\00\00\04\00\00\00\0b\00\00\00\0c\00\00\00\0d\00\00\00\0a\00\00\00\04\00\00\00\04\00\00\00\0e\00\00\00already borrowedalready mutably borrowedassertion failed: `(left == right)`\0a  left: ``,\0a right: ``\00\00P\00\10\00-\00\00\00}\00\10\00\0c\00\00\00\89\00\10\00\01\00\00\00\0f\00\00\00\00\00\00\00\01\00\00\00\10\00\00\00`: \00P\00\10\00-\00\00\00}\00\10\00\0c\00\00\00\b4\00\10\00\03\00\00\00called `Option::unwrap()` on a `None` valuesrc/libcore/option.rs\d0\00\10\00+\00\00\00\fb\00\10\00\15\00\00\00[\01\00\00\15\00\00\00: \00\00\a4\00\10\00\00\00\00\00(\01\10\00\02\00\00\00src/libcore/result.rs\00\00\00<\01\10\00\15\00\00\00\e7\03\00\00\05\00\00\00called `Result::unwrap()` on an `Err` valuesrc/liballoc/raw_vec.rsTried to shrink to a larger capacity\00\00\a6\01\10\00$\00\00\00\8f\01\10\00\17\00\00\00@\02\00\00\09\00\00\00\0a\00\00\00\04\00\00\00\04\00\00\00\07\00\00\00src/libstd/thread/mod.rs\f4\01\10\00\18\00\00\00\8c\03\00\00\13\00\00\00inconsistent park state\00\02\00\00\00park state changed unexpectedly\008\02\10\00\1f\00\00\00\f4\01\10\00\18\00\00\00\89\03\00\00\0d\00\00\00\f4\01\10\00\18\00\00\00\22\04\00\00\11\00\00\00failed to generate unique thread ID: bitspace exhaustedthread name may not contain interior null bytes\00\00\f4\01\10\00\18\00\00\00\97\04\00\00\12\00\00\00inconsistent state in unparksrc/libstd/sync/condvar.rs\00\00\14\03\10\00\1a\00\00\00H\02\00\00\12\00\00\00attempted to use a condition variable with two mutexes\00\00\0a\00\00\00\04\00\00\00\04\00\00\00\11\00\00\00\12\00\00\00src/libstd/sync/once.rs\00\8c\03\10\00\17\00\00\00\8e\01\00\00\15\00\00\00assertion failed: state & STATE_MASK == RUNNING\00\8c\03\10\00\17\00\00\00j\01\00\00\15\00\00\00Once instance has previously been poisoned\00\00\8c\03\10\00\17\00\00\00\c0\01\00\00\09\00\00\00src/libstd/sys_common/at_exit_imp.rs0\04\10\00$\00\00\001\00\00\00\0d\00\00\00assertion failed: queue != DONEPoisonError { inner: .. }src/libstd/sys_common/thread_info.rs\9c\04\10\00$\00\00\00%\00\00\00\1a\00\00\00assertion failed: c.borrow().is_none()\00\00\13\00\00\00\10\00\00\00\04\00\00\00\14\00\00\00\15\00\00\00\16\00\00\00\0c\00\00\00\04\00\00\00\17\00\00\00\18\00\00\00\08\00\00\00\04\00\00\00\19\00\00\00\1a\00\00\00\18\00\00\00\08\00\00\00\04\00\00\00\1b\00\00\00\0f\00\00\00\00\00\00\00\01\00\00\00\1c\00\00\00NulError\0a\00\00\00\04\00\00\00\04\00\00\00\1d\00\00\00src/libstd/sys/wasm/condvar.rs\00\00h\05\10\00\1e\00\00\00\17\00\00\00\09\00\00\00can't block with web assemblysrc/libstd/sys/wasm/mutex.rs\00\00\00\b5\05\10\00\1c\00\00\00\16\00\00\00\09\00\00\00cannot recursively acquire mutexsrc/liballoc/raw_vec.rscapacity overflow\1b\06\10\00\11\00\00\00\04\06\10\00\17\00\00\00\ea\02\00\00\05\00\00\00`\00\00\00..\00\00H\06\10\00\02\00\00\00BorrowErrorBorrowMutError\00\00\00#\00\00\00\00\00\00\00\01\00\00\00$\00\00\00index out of bounds: the len is  but the index is \00\00\80\06\10\00 \00\00\00\a0\06\10\00\12\00\00\00called `Option::unwrap()` on a `None` valuesrc/libcore/option.rs\c4\06\10\00+\00\00\00\ef\06\10\00\15\00\00\00[\01\00\00\15\00\00\00src/libcore/slice/mod.rsindex  out of range for slice of length 4\07\10\00\06\00\00\00:\07\10\00\22\00\00\00\1c\07\10\00\18\00\00\00\09\0a\00\00\05\00\00\00slice index starts at  but ends at \00|\07\10\00\16\00\00\00\92\07\10\00\0d\00\00\00\1c\07\10\00\18\00\00\00\0f\0a\00\00\05\00\00\00src/libcore/str/mod.rs[...]byte index  is out of bounds of `\db\07\10\00\0b\00\00\00\e6\07\10\00\16\00\00\00D\06\10\00\01\00\00\00\c0\07\10\00\16\00\00\00\e1\07\00\00\09\00\00\00begin <= end ( <= ) when slicing `\00\00$\08\10\00\0e\00\00\002\08\10\00\04\00\00\006\08\10\00\10\00\00\00D\06\10\00\01\00\00\00\c0\07\10\00\16\00\00\00\e5\07\00\00\05\00\00\00 is not a char boundary; it is inside  (bytes ) of `\db\07\10\00\0b\00\00\00x\08\10\00&\00\00\00\9e\08\10\00\08\00\00\00\a6\08\10\00\06\00\00\00D\06\10\00\01\00\00\00\c0\07\10\00\16\00\00\00\f2\07\00\00\05\00\00\000x00010203040506070809101112131415161718192021222324252627282930313233343536373839404142434445464748495051525354555657585960616263646566676869707172737475767778798081828384858687888990919293949596979899\00\00%\00\00\00\0c\00\00\00\04\00\00\00&\00\00\00'\00\00\00(\00\00\00    ,\0a, (\0a(,)\0a[])\00\00\00\04\00\00\00\04\00\00\00*\00\00\00+\00\00\00,\00\00\00src/libcore/fmt/mod.rs\00\00\f0\09\10\00\16\00\00\00H\04\00\00(\00\00\00\f0\09\10\00\16\00\00\00T\04\00\00\11\00\00\00\00\00\00\00\00\00\00\00src/libcore/unicode/bool_trie.rs0\0a\10\00 \00\00\00'\00\00\00\19\00\00\000\0a\10\00 \00\00\00(\00\00\00 \00\00\000\0a\10\00 \00\00\00*\00\00\00\19\00\00\000\0a\10\00 \00\00\00+\00\00\00\18\00\00\000\0a\10\00 \00\00\00,\00\00\00 \00\00\00\00\01\03\05\05\06\06\03\07\06\08\08\09\11\0a\1c\0b\19\0c\14\0d\12\0e\16\0f\04\10\03\12\12\13\09\16\01\17\05\18\02\19\03\1a\07\1c\02\1d\01\1f\16 \03+\06,\02-\0b.\010\031\022\02\a9\02\aa\04\ab\08\fa\02\fb\05\fd\04\fe\03\ff\09\adxy\8b\8d\a20WX\8b\8c\90\1c\1d\dd\0e\0fKL\fb\fc./?\5c]_\b5\e2\84\8d\8e\91\92\a9\b1\ba\bb\c5\c6\c9\ca\de\e4\e5\ff\00\04\11\12)147:;=IJ]\84\8e\92\a9\b1\b4\ba\bb\c6\ca\ce\cf\e4\e5\00\04\0d\0e\11\12)14:;EFIJ^de\84\91\9b\9d\c9\ce\cf\0d\11)EIWde\8d\91\a9\b4\ba\bb\c5\c9\df\e4\e5\f0\04\0d\11EIde\80\81\84\b2\bc\be\bf\d5\d7\f0\f1\83\85\86\89\8b\8c\98\a0\a4\a6\a8\a9\ac\ba\be\bf\c5\c7\ce\cf\da\dbH\98\bd\cd\c6\ce\cfINOWY^_\89\8e\8f\b1\b6\b7\bf\c1\c6\c7\d7\11\16\17[\5c\f6\f7\fe\ff\80\0dmq\de\df\0e\0f\1fno\1c\1d_}~\ae\af\bb\bc\fa\16\17\1e\1fFGNOXZ\5c^~\7f\b5\c5\d4\d5\dc\f0\f1\f5rs\8ftu\96\97\c9\ff/_&./\a7\af\b7\bf\c7\cf\d7\df\9a@\97\980\8f\1f\ff\ce\ffNOZ[\07\08\0f\10'/\ee\efno7=?BE\90\91\fe\ffSgu\c8\c9\d0\d1\d8\d9\e7\fe\ff\00 _\22\82\df\04\82D\08\1b\04\06\11\81\ac\0e\80\ab5\1e\15\80\e0\03\19\08\01\04/\044\04\07\03\01\07\06\07\11\0aP\0f\12\07U\08\02\04\1c\0a\09\03\08\03\07\03\02\03\03\03\0c\04\05\03\0b\06\01\0e\15\05:\03\11\07\06\05\10\08V\07\02\07\15\0dP\04C\03-\03\01\04\11\06\0f\0c:\04\1d%\0d\06L m\04j%\80\c8\05\82\b0\03\1a\06\82\fd\03Y\07\15\0b\17\09\14\0c\14\0cj\06\0a\06\1a\06Y\07+\05F\0a,\04\0c\04\01\031\0b,\04\1a\06\0b\03\80\ac\06\0a\06\1fAL\04-\03t\08<\03\0f\03<\078\08*\06\82\ff\11\18\08/\11-\03 \10!\0f\80\8c\04\82\97\19\0b\15\88\94\05/\05;\07\02\0e\18\09\80\af1t\0c\80\d6\1a\0c\05\80\ff\05\80\b6\05$\0c\9b\c6\0a\d20\10\84\8d\037\09\81\5c\14\80\b8\08\80\ba=5\04\0a\068\08F\08\0c\06t\0b\1e\03Z\04Y\09\80\83\18\1c\0a\16\09F\0a\80\8a\06\ab\a4\0c\17\041\a1\04\81\da&\07\0c\05\05\80\a5\11\81m\10x(*\06L\04\80\8d\04\80\be\03\1b\03\0f\0d\00\06\01\01\03\01\04\02\08\08\09\02\0a\05\0b\02\10\01\11\04\12\05\13\11\14\02\15\02\17\02\1a\02\1c\05\1d\08$\01j\03k\02\bc\02\d1\02\d4\0c\d5\09\d6\02\d7\02\da\01\e0\05\e8\02\ee \f0\04\f9\04\0c';>NO\8f\9e\9e\9f\06\07\096=>V\f3\d0\d1\04\14\1867VW\bd5\ce\cf\e0\12\87\89\8e\9e\04\0d\0e\11\12)14:EFIJNOdeZ\5c\b6\b7\1b\1c\84\85\097\90\91\a8\07\0a;>fi\8f\92o_\ee\efZb\9a\9b'(U\9d\a0\a1\a3\a4\a7\a8\ad\ba\bc\c4\06\0b\0c\15\1d:?EQ\a6\a7\cc\cd\a0\07\19\1a\22%\c5\c6\04 #%&(38:HJLPSUVXZ\5c^`cefksx}\7f\8a\a4\aa\af\b0\c0\d0?qr{^\22{\05\03\04-\03e\04\01/.\80\82\1d\031\0f\1c\04$\09\1e\05+\05D\04\0e*\80\aa\06$\04$\04(\084\0b\01\80\90\817\09\16\0a\08\80\989\03c\08\090\16\05!\03\1b\05\01@8\04K\05/\04\0a\07\09\07@ '\04\0c\096\03:\05\1a\07\04\0c\07PI73\0d3\07.\08\0a\81&\1f\80\81(\08*\80\a6N\04\1e\0fC\0e\19\07\0a\06G\09'\09u\0b?A*\06;\05\0a\06Q\06\01\05\10\03\05\80\8b_!H\08\0a\80\a6^\22E\0b\0a\06\0d\138\08\0a6,\04\10\80\c0<dS\0c\01\81\00H\08S\1d9\81\07F\0a\1d\03GI7\03\0e\08\0a\069\07\0a\816\19\81\07\83\9afu\0b\80\c4\8a\bc\84/\8f\d1\82G\a1\b9\829\07*\04\02`&\0aF\0a(\05\13\82\b0[eE\0b/\10\11@\02\1e\97\f2\0e\82\f3\a5\0d\81\1fQ\81\8c\89\04k\05\0d\03\09\07\10\93`\80\f6\0as\08n\17F\80\9a\14\0cW\09\19\80\87\81G\03\85B\0f\15\85P+\87\d5\80\d7)K\05\0a\04\02\83\11D\81K<\06\01\04U\05\1b4\02\81\0e,\04d\0cV\0a\0d\03\5c\04=9\1d\0d,\04\09\07\02\0e\06\80\9a\83\d5\0b\0d\03\0a\06t\0cY'\0c\048\08\0a\06(\08\1eR\0c\04g\03)\0d\0a\06\03\0d0`\0e\85\92\00\00\c0\fb\ef>\00\00\00\00\00\0e\00\00\00\00\00\00\00\00\00\00\00\00\00\00\f8\ff\fb\ff\ff\ff\07\00\00\00\00\00\00\14\fe!\fe\00\0c\00\00\00\02\00\00\00\00\00\00P\1e \80\00\0c\00\00@\06\00\00\00\00\00\00\10\869\02\00\00\00#\00\be!\00\00\0c\00\00\fc\02\00\00\00\00\00\00\d0\1e \c0\00\0c\00\00\00\04\00\00\00\00\00\00@\01 \80\00\00\00\00\00\11\00\00\00\00\00\00\c0\c1=`\00\0c\00\00\00\02\00\00\00\00\00\00\90D0`\00\0c\00\00\00\03\00\00\00\00\00\00X\1e \80\00\0c\00\00\00\00\84\5c\80\00\00\00\00\00\00\00\00\00\00\f2\07\80\7f\00\00\00\00\00\00\00\00\00\00\00\00\f2\1b\00?\00\00\00\00\00\00\00\00\00\03\00\00\a0\02\00\00\00\00\00\00\fe\7f\df\e0\ff\fe\ff\ff\ff\1f@\00\00\00\00\00\00\00\00\00\00\00\00\e0\fdf\00\00\00\c3\01\00\1e\00d \00 \00\00\00\00\00\00\00\e0\00\00\00\00\00\00\1c\00\00\00\1c\00\00\00\0c\00\00\00\0c\00\00\00\00\00\00\00\b0?@\fe\0f \00\00\00\00\008\00\00\00\00\00\00`\00\00\00\00\02\00\00\00\00\00\00\87\01\04\0e\00\00\80\09\00\00\00\00\00\00@\7f\e5\1f\f8\9f\00\00\00\00\00\00\ff\7f\0f\00\00\00\00\00\d0\17\04\00\00\00\00\f8\0f\00\03\00\00\00<;\00\00\00\00\00\00@\a3\03\00\00\00\00\00\00\f0\cf\00\00\00\f7\ff\fd!\10\03\ff\ff\ff\ff\ff\ff\ff\fb\00\10\00\00\00\00\00\00\00\00\ff\ff\ff\ff\01\00\00\00\00\00\00\80\03\00\00\00\00\00\00\00\00\80\00\00\00\00\ff\ff\ff\ff\00\00\00\00\00\fc\00\00\00\00\00\06\00\00\00\00\00\00\00\00\00\80\f7?\00\00\00\c0\00\00\00\00\00\00\00\00\00\00\03\00D\08\00\00`\00\00\000\00\00\00\ff\ff\03\80\00\00\00\00\c0?\00\00\80\ff\03\00\00\00\00\00\07\00\00\00\00\00\c8\13\00\00\00\00 \00\00\00\00\00\00\00\00~f\00\08\10\00\00\00\00\00\10\00\00\00\00\00\00\9d\c1\02\00\00\00\000@\00\00\00\00\00 !\00\00\00\00\00@\00\00\00\00\ff\ff\00\00\ff\ff\00\00\00\00\00\00\00\00\00\01\00\00\00\02\00\03\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\04\00\00\05\00\00\00\00\00\00\00\00\06\00\00\00\00\00\00\00\00\07\00\00\08\09\0a\00\0b\0c\0d\0e\0f\00\00\10\11\12\00\00\13\14\15\16\00\00\17\18\19\1a\1b\00\1c\00\00\00\1d\00\00\00\00\00\00\00\1e\1f \00\00\00\00\00!\00\22\00#$%\00\00\00\00&\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00'(\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00)\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00*\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00+,\00\00-\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00./0\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\001\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\002\003\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\0045\00\005556\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00 \00\00\00\00\01\00\00\00\00\00\00\00\00\00\c0\07n\f0\00\00\00\00\00\87\00\00\00\00`\00\00\00\00\00\00\00\f0\00\00\00\c0\ff\01\00\00\00\00\00\02\00\00\00\00\00\00\ff\7f\00\00\00\00\00\00\80\03\00\00\00\00\00x\06\07\00\00\00\80\ef\1f\00\00\00\00\00\00\00\08\00\03\00\00\00\00\00\c0\7f\00\1e\00\00\00\00\00\00\00\00\00\00\00\80\d3@\00\00\00\80\f8\07\00\00\03\00\00\00\00\00\00X\01\00\80\00\c0\1f\1f\00\00\00\00\00\00\00\00\ff\5c\00\00@\00\00\00\00\00\00\00\00\00\00\f9\a5\0d\00\00\00\00\00\00\00\00\00\00\00\00\80<\b0\01\00\000\00\00\00\00\00\00\00\00\00\00\f8\a7\01\00\00\00\00\00\00\00\00\00\00\00\00(\bf\00\00\00\00\e0\bc\0f\00\00\00\00\00\00\00\80\ff\06\fe\07\00\00\00\00\f8y\80\00~\0e\00\00\00\00\00\fc\7f\03\00\00\00\00\00\00\00\00\00\00\7f\bf\00\00\fc\ff\ff\fcm\00\00\00\00\00\00\00~\b4\bf\00\00\00\00\00\00\00\00\00\a3\00\00\00\00\00\00\00\00\00\00\00\18\00\00\00\00\00\00\00\1f\00\00\00\00\00\00\00\7f\00\00\80\07\00\00\00\00\00\00\00\00`\00\00\00\00\00\00\00\00\a0\c3\07\f8\e7\0f\00\00\00<\00\00\1c\00\00\00\00\00\00\00\ff\ff\ff\ff\ff\ff\7f\f8\ff\ff\ff\ff\ff\1f \00\10\00\00\f8\fe\ff\00\00\7f\ff\ff\f9\db\07\00\00\00\00\7f\00\00\00\00\00\f0\07\00\00\00\00\00\00\00\00\00\00\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\ff\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\f8\03\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\fe\ff\ff\ff\ff\bf\b6\00\00\00\00\00\00\00\00\00\ff\07\00\00\00\00\00\f8\ff\ff\00\00\01\00\00\00\00\00\00\00\00\00\00\00\c0\9f\9f=\00\00\00\00\02\00\00\00\ff\ff\ff\07\00\00\00\00\00\00\00\00\00\00\c0\ff\01\00\00\00\00\00\00\f8\0f \b8\0f\10\00J\00\00\00\08\12\10\00\00\02\00\00\08\14\10\007\00\00\00\00\01\02\03\04\05\06\07\08\09\08\0a\0b\0c\0d\0e\0f\10\11\12\13\14\02\15\16\17\18\19\1a\1b\1c\1d\1e\1f \02\02\02\02\02\02\02\02\02\02!\02\02\02\02\02\02\02\02\02\02\02\02\02\02\22#$%&\02'\02(\02\02\02)*+\02,-./0\02\021\02\02\022\02\02\02\02\02\02\02\023\02\024\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\025\026\027\02\02\02\02\02\02\02\028\029\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02:;<\02\02\02\02=\02\02>?@ABCDEF\02\02\02G\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02H\02\02\02\02\02\02\02\02\02\02\02I\02\02\02\02\02;\02\00\01\02\02\02\02\03\02\02\02\02\04\02\05\06\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\07\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02\02.\00\00\00\04\00\00\00\04\00\00\00/\00\00\00/\00\00\000\00\00\00")
  (data (;1;) (i32.const 1055696) "\01\00\00\00\00\00\00\00")
  (data (;2;) (i32.const 1055704) "\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"))
